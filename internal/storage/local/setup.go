package local

import "os"

func MustSetup(storageDir string) string {
	if _, err := os.Stat(storageDir); err != nil {
		if err = os.MkdirAll(storageDir, os.ModePerm); err != nil {
			panic("failed to init data dir: " + err.Error())
		}
	}

	return storageDir
}
