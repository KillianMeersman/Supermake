package util

import (
	"errors"
	"os"
)

func IsFileOrFolder(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}
