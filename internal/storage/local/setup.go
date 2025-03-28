package local

import (
	"os"
	"path"
)

func MustSetup(storageDir string, dirs []string) string {
	if _, err := os.Stat(storageDir); err != nil {
		if err = os.Mkdir(storageDir, os.ModePerm); err != nil {
			panic("failed to init data root directory: " + err.Error())
		}
	}

	for _, dir := range dirs {
		dirPath := path.Join(storageDir, dir)

		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			panic("failed to init data custom directory: " + err.Error())
		}
	}

	return storageDir
}
