package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"fyne.io/fyne/v2"
)

var createNewFolderError = errors.New("Не удалось создать папку для фотографий")

func GetJpegPaths(files []fyne.URI) []string {
	var jpegPaths []string

	for _, file := range files {
		if file.MimeType() == "image/jpeg" {
			jpegPaths = append(jpegPaths, file.Path())
		}
	}

	return jpegPaths
}

func CreateNewFolder() (string, error) {
	currentDirName := createNewFolderName()
	if err := os.Mkdir(currentDirName, os.ModePerm); err != nil {
		fmt.Println(err)
		return "", errors.Join(createNewFolderError, err)
	}

	return currentDirName, nil
}

func createNewFolderName() string {
	currentTime := time.Now().Format("2006-01-02_15-04-05")

	return filepath.Join(createDesktopFolderName(), currentTime)
}

func createDesktopFolderName() string {
	homeDir, _ := os.UserHomeDir()

	return filepath.Join(homeDir, "Desktop")
}
