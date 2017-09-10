package storage

import (
	"os"
	"path/filepath"
)

type Storage struct {
	Schema string
}

var fileStorage = Storage{
	Schema: "file://",
}

//IsExit simple func check file/directory
func IsExit(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func Abs(path string) (string, error) {
	fullPath, err := filepath.Abs(path)
	return fileStorage.Schema + ToSlash(fullPath), err
}

func ToSlash(path string) string {
	return filepath.ToSlash(filepath.Clean(path))
}
