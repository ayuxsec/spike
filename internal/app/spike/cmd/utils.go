package cmd

import (
	"bufio"
	"fmt"
	"os"
	"spike/pkg/config"
	"strings"
)

func ensureAppDirExists() error {
	return os.MkdirAll(config.DefaultAppDir, 0755)
}

// fileToSlice reads a file and returns a slice of strings, each representing a line
func fileToSlice(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var lines []string
	bufioScanner := bufio.NewScanner(file)
	for bufioScanner.Scan() {
		line := strings.TrimSpace(bufioScanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := bufioScanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return lines, nil
}
