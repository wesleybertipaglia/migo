package cmd

import (
	"fmt"
	"log"
	"migo/db"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all migrations and their status",
	Run: func(cmd *cobra.Command, args []string) {
		var rootDir string = "./migo"

		db.Init(rootDir)
		defer db.DB.Close()

		fmt.Println("✅ Applied Migrations:")
		rows, err := db.DB.Query("SELECT timestamp, name FROM migrations_applied ORDER BY timestamp ASC")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var timestamp, name string
			if err := rows.Scan(&timestamp, &name); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("- %s %s (Applied)\n", timestamp, name)
		}

		fmt.Println("\n⚠️ Pending Migrations:")
		rows, err = db.DB.Query("SELECT timestamp, name FROM migrations_pending ORDER BY timestamp ASC")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var timestamp, name string
			if err := rows.Scan(&timestamp, &name); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("- %s %s (Pending)\n", timestamp, name)
		}
	},
}
