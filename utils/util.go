package utils

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/Saikatdeb12/TodoApp/internal/middleware"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string{
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}

func GetUserID(ctx context.Context) (uuid.UUID, error) {
	userId, ok := ctx.Value(middlewares.UserIDkey).(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("Unauthorized")
	}

	return userId,nil
}
