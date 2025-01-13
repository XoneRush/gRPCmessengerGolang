package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Подключение к базе данных
func Connect(props Properties) *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		props.Host, props.Port, props.User, props.Password, props.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to", props.Dbname, "! At port:", props.Port)
	return db
}
