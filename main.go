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

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
