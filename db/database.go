package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Driver SQLite3
)

var Database *sql.DB

// InitDatabase inicializa o banco de dados SQLite
func InitDatabase() {
	var err error
	Database, err = sql.Open("sqlite3", "./db/sitevila.db") // Nome do banco de dados
	if err != nil {
		log.Fatalf("Erro ao abrir o banco de dados: %v", err)
	}

	createTables()
}

// createTables cria as tabelas necessárias no banco de dados
func createTables() {
	_, err := Database.Exec(`
	CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		type TEXT NOT NULL, -- 'deposit' ou 'expense'
		amount REAL NOT NULL,
		date TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS residents (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		contributed BOOLEAN NOT NULL
	);
	`)
	if err != nil {
		log.Fatalf("Erro ao criar tabelas: %v", err)
	}
}
