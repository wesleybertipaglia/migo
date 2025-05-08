package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateProjectStructure(baseDir string) error {
	dirs := []string{"migrations", "logs", "state"}
	for _, dir := range dirs {
		fullPath := filepath.Join(baseDir, dir)
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			return fmt.Errorf("failed to create %s: %w", fullPath, err)
		}
	}
	return nil
}
