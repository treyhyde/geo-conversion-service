package internal

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// ConvertToMBTile converts a georeferenced file to an MBTile pack using ogr2ogr
func ConvertToMBTile(inputFilePath string) (string, error) {
	outputFilePath := strings.TrimSuffix(inputFilePath, filepath.Ext(inputFilePath)) + ".mbtiles"
	cmd := exec.Command("ogr2ogr", "-f", "MBTiles", outputFilePath, inputFilePath)

	// Ensure the command is constructed safely to prevent command injection
	if err := validateCommand(cmd); err != nil {
		return "", fmt.Errorf("invalid command: %v", err)
	}

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to convert file: %v", err)
	}

	return outputFilePath, nil
}

// validateCommand ensures the command is constructed safely to prevent command injection
func validateCommand(cmd *exec.Cmd) error {
	// Add validation logic to ensure the command is safe
	// For simplicity, we'll just check that the command is "ogr2ogr"
	if cmd.Path != "ogr2ogr" {
		return fmt.Errorf("invalid command path: %s", cmd.Path)
	}

	// Check that the arguments are safe
	for _, arg := range cmd.Args {
		if strings.Contains(arg, ";") || strings.Contains(arg, "&") || strings.Contains(arg, "|") {
			return fmt.Errorf("invalid command argument: %s", arg)
		}
	}

	return nil
}
