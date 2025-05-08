package main

import (
	"log"
	"os"

	"migo/cmd"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "migo",
		Short: "Migo is a migration tool for managing database migrations",
	}

	rootCmd.AddCommand(cmd.InitCmd)
	rootCmd.AddCommand(cmd.AddCmd)
	rootCmd.AddCommand(cmd.UpdateCmd)
	rootCmd.AddCommand(cmd.RollbackCmd)
	rootCmd.AddCommand(cmd.ListCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
