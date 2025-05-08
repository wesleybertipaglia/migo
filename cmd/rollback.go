package cmd

import (
	"bufio"
	"fmt"
	"log"
	"migo/db"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var RollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback the last applied migration",
	Run: func(cmd *cobra.Command, args []string) {
		var rootDir string = "./migo"

		migrationsPath := filepath.Join(rootDir, "migrations")
		if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
			fmt.Println("❌ Invalid directory: migo folder not found.")
			return
		}

		db.Init(rootDir)

		row := db.DB.QueryRow("SELECT timestamp, name FROM migrations_applied ORDER BY id DESC LIMIT 1")
		var timestamp, name string
		if err := row.Scan(&timestamp, &name); err != nil {
			fmt.Println("⚠️ No migrations to rollback.")
			return
		}

		filePath := filepath.Join(migrationsPath, fmt.Sprintf("%s_%s.sql", timestamp, name))
		downSQL, err := extractDownSQL(filePath)
		if err != nil {
			log.Fatalf("Error reading migration file: %v", err)
		}

		if downSQL == "" {
			fmt.Println("⚠️ No DOWN SQL section found.")
			return
		}

		fmt.Printf("⏪ Rolling back migration: %s\n", name)
		if _, err := db.DB.Exec(downSQL); err != nil {
			log.Fatalf("❌ Failed to rollback migration: %v", err)
		}

		_, err = db.DB.Exec("DELETE FROM migrations_applied WHERE timestamp = ?", timestamp)
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.DB.Exec("INSERT INTO migrations_pending (timestamp, name, created_at) VALUES (?, ?, datetime('now'))", timestamp, name)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("✅ Rolled back migration: %s\n", name)
	},
}

func extractDownSQL(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var downSection bool
	var builder strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		if downSection {
			builder.WriteString(line + "\n")
		}
		if strings.HasPrefix(line, "-- DOWN") {
			downSection = true
		}
	}

	return builder.String(), scanner.Err()
}
