package cmd

import (
	"fmt"
	"log"
	"migo/db"
	"migo/validations"
	"os"
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

		validations.ValidateDirectory(rootDir)
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
		defer db.DB.Close()

		tx, err := db.DB.Begin()
		if err != nil {
			log.Fatalf("❌ Failed to begin transaction: %v", err)
		}
		defer tx.Commit()

		_, err = tx.Exec("INSERT INTO migrations_pending (timestamp, name, created_at) VALUES (?, ?, ?)",
			timestamp, migrationName, time.Now().Format(time.RFC3339))
		if err != nil {
			tx.Rollback()
			log.Fatalf("❌ Failed to insert new migration into pending migrations: %v", err)
		}

		fmt.Printf("✅ Migration '%s' created at: %s\n", migrationName, fileName)
	},
}

func init() {
	AddCmd.Flags().StringVarP(&migrationName, "name", "n", "", "Migration name")
}
