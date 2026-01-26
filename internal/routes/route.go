package routes

import (
	"github.com/Saikatdeb12/TodoApp/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux{
	r := chi.NewRouter()
	r.Post("/auth/register", handlers.Register)
	r.Post("/auth/login", handlers.Login)

	return r
}
