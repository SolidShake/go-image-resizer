package file

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"fyne.io/fyne/v2"
)

func GetJpegPaths(files []fyne.URI) []string {
	var jpegPaths []string

	for _, file := range files {
		if file.MimeType() == "image/jpeg" {
			jpegPaths = append(jpegPaths, file.Path())
		}
	}

	return jpegPaths
}

func CreateFolder(userDir string) (string, error) {
	currentDirName := createFolderName(userDir)
	if err := os.Mkdir(currentDirName, os.ModePerm); err != nil {
		fmt.Println(err)
		return "", err
	}

	return currentDirName, nil
}

func createFolderName(userDir string) string {
	currentTime := time.Now().Format("2006-01-02_15-04-05")

	return filepath.Join(userDir, currentTime)
}
