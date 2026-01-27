package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID   uuid.UUID `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"-" db:"password"`
	Todos    []Todo
}


type Todo struct {
	TodoID    uuid.UUID `json:"id" db:"id"`
	UserID    uuid.UUID `json:"-" db:"user_id"`
	Title     string    `json:"title" db:"title"`
	Body      string    `json:"body" db:"body"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	ValidTill time.Time `json:"validTill" db:"valid_till"`
	Complete  bool      `json:"complete" db:"complete"`
}

type Session struct {
	SessionID uuid.UUID `db:"session_id"`
	UserID    uuid.UUID `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	ExpiresAt time.Time `db:"expires_at"`
}

// run the container and run commands:
// docker exec -it 2b9b7857fb42 psql -U todo_user -d todo_db
