package database

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func Connect(){
	connStr := "postgres://local:local@localhost:5432/todo_db?sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if(err != nil){
		panic("DB open failed!")
	}

	if err=DB.Ping(); err!=nil{
		panic("DB ping error")
	}

	fmt.Println("Connected to PostgreSQL")
}
