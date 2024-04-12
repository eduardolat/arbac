package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

func ReadGlobFiles(patterns []string) ([][]byte, error) {
	paths := []string{}
	for _, pattern := range patterns {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			return nil, fmt.Errorf("error globbing pattern %s: %w", pattern, err)
		}
		for _, match := range matches {
			if !slices.Contains(paths, match) {
				paths = append(paths, match)
			}
		}
	}

	var dataSlice [][]byte
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("error reading file %s: %w", path, err)
		}
		dataSlice = append(dataSlice, data)
	}

	return dataSlice, nil
}
