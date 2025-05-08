package cmd

import (
	"fmt"
	"log"
	"migo/db"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var migrationName string

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new migration",
	Run: func(cmd *cobra.Command, args []string) {
		var rootDir string = "./migo"

		if migrationName == "" {
			fmt.Println("❌ Missing migration name.")
			return
		}

		migrationsPath := filepath.Join(rootDir, "migrations")
		if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
			fmt.Println("❌ Invalid directory: migo folder not found.")
			return
		}

		timestamp := time.Now().Format("20060102150405")
		fileName := fmt.Sprintf("%s/migrations/%s_%s.sql", rootDir, timestamp, migrationName)

		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		content := fmt.Sprintf("-- Migration: %s\n-- Created at: %s\n\n-- UP\n\n-- DOWN\n", migrationName, timestamp)
		if _, err = file.WriteString(content); err != nil {
			log.Fatal(err)
		}

		db.Init(rootDir)
		_, err = db.DB.Exec("INSERT INTO migrations_pending (timestamp, name, created_at) VALUES (?, ?, ?)",
			timestamp, migrationName, time.Now().Format(time.RFC3339))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("✅ Migration '%s' created at: %s\n", migrationName, fileName)
	},
}

func init() {
	AddCmd.Flags().StringVarP(&migrationName, "name", "n", "", "Migration name")
}
