package db

import (
	"database/sql"
	"log"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init(dir string) {
	var err error
	dbPath := filepath.Join(dir, "state", "migo.db")

	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	createTables := `
	CREATE TABLE IF NOT EXISTS migrations_applied (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp TEXT NOT NULL,
		name TEXT NOT NULL,
		applied_at TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS migrations_pending (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp TEXT NOT NULL,
		name TEXT NOT NULL,
		created_at TEXT NOT NULL
	);
	`

	if _, err := DB.Exec(createTables); err != nil {
		log.Fatal(err)
	}
}
