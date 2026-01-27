package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Saikatdeb12/TodoApp/database"
	"github.com/Saikatdeb12/TodoApp/internal/models"
	"github.com/Saikatdeb12/TodoApp/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type CreateTodoRequest struct {
	Title string `json:"title"`
	Body string `json:"body"`
	ValidTill time.Time `json:"validTill"`
}

func itoa(i int) string {
	return fmt.Sprintf("%d", i)
}

func CreateTodo(w http.ResponseWriter, r *http.Request){
	var req CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	userId, err := utils.GetUserID(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	query := `
		INSERT INTO todos (user_id, title, body, valid_till)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`

	var todo models.Todo
	err = database.DB.QueryRow(
		query, userId, req.Title, req.Body, req.ValidTill,
	).Scan(&todo.TodoID, &todo.CreatedAt)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todo.Title=req.Title
	todo.Body=req.Body
	todo.ValidTill=req.ValidTill
		
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)

}

func GetTodos(w http.ResponseWriter, r *http.Request){
	userId, err := utils.GetUserID(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	query := `
		SELECT id, title, body, created_at, complete, valid_till
		FROM todos
		WHERE user_id=$1
	`

	args := []interface{}{userId}
	argID := 2
	if c:= r.URL.Query().Get("complete"); c!=""{
		query+=" AND complet=$" + itoa(argID)
		args=append(args, c=="true")
		argID++
	}

	if from := r.URL.Query().Get("from"); from != "" {
		query+= " AND created_at >=$" + itoa(argID)
		args=append(args, from)
		argID++
	}

	if to := r.URL.Query().Get("to"); to !="" {
		query += " AND created_at <= $" + itoa(argID)
		args=append(args, to)
		argID++
	}

	query+=" ORDER BY created_at DESC"

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// always to close the db connection
	defer rows.Close()

	todos := []models.Todo{}
	for rows.Next(){
		var t models.Todo
		err := rows.Scan(
			&t.TodoID,
			&t.Title,
			&t.Body,
			&t.CreatedAt,
			&t.Complete,
			&t.ValidTill,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		todos=append(todos, t)
	}
	json.NewEncoder(w).Encode(todos)
}

func GetTodoByID (w http.ResponseWriter, r *http.Request){
	userID, err := utils.GetUserID(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	todoID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid todo id", http.StatusBadRequest)
		return
	}

	query := `
		SELECT id, title, body, created_at, complete, valid_till
		FROM todos
		WHERE id=$1 AND user_id=$2
	`

	var todo models.Todo

	err = database.DB.QueryRow(query, todoID, userID).Scan(
		&todo.TodoID,
		&todo.Title,
		&todo.Body,
		&todo.CreatedAt,
		&todo.Complete,
		&todo.ValidTill,
	)

	if err == sql.ErrNoRows {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

