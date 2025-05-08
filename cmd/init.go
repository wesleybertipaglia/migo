package cmd

import (
	"fmt"
	"log"
	"migo/db"
	"migo/utils"
	"path/filepath"

	"github.com/spf13/cobra"
)

var dir string

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a migration project in a directory (default: current directory)",
	Run: func(cmd *cobra.Command, args []string) {
		if dir == "" {
			dir = "."
		}

		dir := filepath.Join(dir, "migo")

		fmt.Printf("🚀 Initializing project at: %s\n", dir)

		if err := utils.CreateProjectStructure(dir); err != nil {
			log.Fatalf("❌ Failed to create project structure: %v", err)
		}

		db.Init(dir)

		fmt.Println("✅ Project initialized.")
	},
}

func init() {
	InitCmd.Flags().StringVarP(&dir, "project", "p", "", "Directory to create the project in (default: current directory)")
}
