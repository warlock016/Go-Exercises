package output

import (
	"errors"
	"os"
)

func DumpToCache(filePath string, byteData []byte) error {
	_, err := os.Open(filePath)
	if os.IsExist(err) {
		return errors.New("file already exists")
	}
	return os.WriteFile(filePath, byteData, 0o644)
}
