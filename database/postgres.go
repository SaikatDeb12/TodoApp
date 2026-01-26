package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/Saikatdeb12/TodoApp/utils"
)

var DB *sql.DB

func Connect(){
	connStr := utils.GoDotEnvVariable("POSTGRESQL_URL")
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
