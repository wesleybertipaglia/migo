package validations

import (
	"fmt"
	"os"
	"path/filepath"
)

func ValidateDirectory(rootDir string) error {
	if rootDir == "" {
		return fmt.Errorf("❌ Missing project directory")
	}

	migrationsPath := filepath.Join(rootDir, "migrations")
	if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
		return fmt.Errorf("❌ Invalid directory: migrations folder not found")
	}

	statePath := filepath.Join(rootDir, "state")
	if _, err := os.Stat(statePath); os.IsNotExist(err) {
		return fmt.Errorf("❌ Invalid directory: state folder not found")
	}

	logsPath := filepath.Join(rootDir, "logs")
	if _, err := os.Stat(logsPath); os.IsNotExist(err) {
		return fmt.Errorf("❌ Invalid directory: logs folder not found")
	}

	return nil
}
