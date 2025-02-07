package test

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yourusername/geo-conversion-service/internal"
)

func TestConvertToMBTile(t *testing.T) {
	// Test valid file conversion
	inputFilePath := "testdata/valid_file.shp"
	expectedOutputFilePath := "testdata/valid_file.mbtiles"
	outputFilePath, err := internal.ConvertToMBTile(inputFilePath)
	assert.NoError(t, err)
	assert.Equal(t, expectedOutputFilePath, outputFilePath)

	// Test invalid file conversion
	inputFilePath = "testdata/invalid_file.txt"
	outputFilePath, err = internal.ConvertToMBTile(inputFilePath)
	assert.Error(t, err)
	assert.Empty(t, outputFilePath)
}

func TestValidateCommand(t *testing.T) {
	// Test valid command
	cmd := exec.Command("ogr2ogr", "-f", "MBTiles", "output.mbtiles", "input.shp")
	err := internal.ValidateCommand(cmd)
	assert.NoError(t, err)

	// Test invalid command path
	cmd = exec.Command("invalid_command", "-f", "MBTiles", "output.mbtiles", "input.shp")
	err = internal.ValidateCommand(cmd)
	assert.Error(t, err)

	// Test invalid command argument
	cmd = exec.Command("ogr2ogr", "-f", "MBTiles", "output.mbtiles", "input.shp; rm -rf /")
	err = internal.ValidateCommand(cmd)
	assert.Error(t, err)
}
