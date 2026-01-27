package routes

import (
	"github.com/Saikatdeb12/TodoApp/internal/handlers"
	"github.com/Saikatdeb12/TodoApp/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux{
	r := chi.NewRouter()
	r.Post("/auth/register", handlers.Register)
	r.Post("/auth/login", handlers.Login)
	r.Post("/auth/logout", handlers.Logout)
	r.Group(func (r chi.Router){
		r.Use(middlewares.Auth)
		r.Get("/todos", handlers.GetTodos)
		r.Get("/todos/{id}", handlers.GetTodoByID)
		r.Post("/todos", handlers.CreateTodo)
	})

	return r
}
