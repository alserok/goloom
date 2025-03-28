package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
	TB = GB * 1024
)

func PrettifySize(size int64) string {
	if size < KB {
		return fmt.Sprintf("%d B", size)
	}

	if size < MB {
		return fmt.Sprintf("%.2f KB", float64(size/KB))
	}

	if size < GB {
		return fmt.Sprintf("%.2f MB", float64(size/MB))
	}

	if size < TB {
		return fmt.Sprintf("%.2f GB", float64(size/GB))
	}

	return fmt.Sprintf("%.2f TB", float64(size/TB))
}

func CountSize(root string) int64 {
	var totalSize int64
	_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			totalSize += info.Size()
		}

		return nil
	})

	return totalSize
}
