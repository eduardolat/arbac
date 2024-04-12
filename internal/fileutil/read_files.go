package fileutil

import (
	"fmt"
	"os"
)

func ReadFiles(paths []string) ([][]byte, error) {
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
