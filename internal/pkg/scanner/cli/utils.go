package cli

import (
	"os"
	"strings"
)

func IsFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	if fileInfo.IsDir() {
		return false
	}
	return true
}

func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func RemoveDuplicatesAndEmptyStrings(slice []string) []string {
	unique := make(map[string]struct{})
	result := []string{}

	for _, item := range slice {
		if _, exists := unique[item]; !exists {
			unique[item] = struct{}{}
			if item != "" {
				result = append(result, item)
			}
		}
	}

	return result
}

func JoinSlice(slice []string) string {
	return strings.Join(slice, "\n")
}

func LinesToSlice(lines string) []string {
	return strings.Split(strings.TrimSpace(lines), "\n")
}
