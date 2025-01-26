package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func test() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Conexão com SQLite bem-sucedida!")
}
