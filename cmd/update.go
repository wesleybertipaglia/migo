package cmd

import (
	"bufio"
	"fmt"
	"log"
	"migo/db"
	"migo/validations"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var mutex sync.Mutex

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Apply all pending migrations",
	Run: func(cmd *cobra.Command, args []string) {
		var rootDir string = "./migo"

		if rootDir == "" {
			fmt.Println("❌ Missing project directory.")
			return
		}

		validations.ValidateDirectory(rootDir)
		migrationsPath := filepath.Join(rootDir, "migrations")

		db.Init(rootDir)
		defer db.DB.Close()

		mutex.Lock()
		defer mutex.Unlock()

		tx, err := db.DB.Begin()
		if err != nil {
			log.Fatalf("❌ Failed to begin transaction: %v", err)
		}

		defer tx.Commit()

		rows, err := tx.Query("SELECT timestamp, name FROM migrations_pending ORDER BY timestamp")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var timestamp, name string
			if err := rows.Scan(&timestamp, &name); err != nil {
				log.Fatal(err)
			}

			filePath := filepath.Join(migrationsPath, fmt.Sprintf("%s_%s.sql", timestamp, name))
			upSQL, err := extractUpSQL(filePath)
			if err != nil {
				log.Fatalf("Error reading migration file %s: %v", filePath, err)
			}

			if upSQL != "" {
				fmt.Printf("🚀 Applying migration: %s\n", name)
				if _, err := tx.Exec(upSQL); err != nil {
					tx.Rollback()
					log.Fatalf("❌ Failed to apply migration %s: %v", name, err)
				}

				_, err = tx.Exec("DELETE FROM migrations_pending WHERE timestamp = ?", timestamp)
				if err != nil {
					tx.Rollback()
					log.Fatal(err)
				}

				_, err = tx.Exec("INSERT INTO migrations_applied (timestamp, name, applied_at) VALUES (?, ?, ?)",
					timestamp, name, time.Now().Format(time.RFC3339))
				if err != nil {
					tx.Rollback()
					log.Fatal(err)
				}

				fmt.Printf("✅ Migration '%s' applied successfully.\n", name)
			}
		}
	},
}

func extractUpSQL(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var upSection bool
	var builder strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "-- DOWN") {
			break
		}
		if upSection {
			builder.WriteString(line + "\n")
		}
		if strings.HasPrefix(line, "-- UP") {
			upSection = true
		}
	}

	return builder.String(), scanner.Err()
}
