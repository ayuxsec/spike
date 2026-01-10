package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func fileToSlice(path string) ([]string, error) {
	var slice []string
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			slice = append(slice, strings.TrimSpace(scanner.Text()))
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("bufio failed to read file: %v", err)
	}
	return slice, nil
}
