package db

import (
	"database/sql"
	"log"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init(dir string) {
	dbPath := filepath.Join(dir, "state", "migo.db")

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec("PRAGMA busy_timeout = 5000"); err != nil {
		log.Fatal(err)
	}

	DB = db

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
