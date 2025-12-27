package input

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type FilePaths []string

func FindFiles(location, fileType string) (*FilePaths, error) {
	info, err := os.Stat(location)
	if err != nil {
		return nil, fmt.Errorf("unexpected path error: %v - %s\n", err, location)
	}

	if !info.IsDir() {
		if strings.Contains(info.Name(), fileType) {
			return &FilePaths{location}, nil
		}
		return nil, fmt.Errorf("location is not dir: %s", location)
	} else {
		f, err := os.ReadDir(location)
		if err != nil {
			return nil, fmt.Errorf("folder read error: %v - %s\n", err, location)
		}
		result := make(FilePaths, 0, 10)
		for _, file := range f {
			if strings.Contains(file.Name(), fileType) {
				result = append(result, path.Join(location, file.Name()))
			}
		}
		return &result, nil
	}
}
