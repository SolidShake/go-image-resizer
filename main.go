package main

import (
	"fmt"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/SolidShake/go-image-resizer/internal/watermark"
)

func main() {
	a := app.New()
	w := a.NewWindow("Watermark Photo Resizer")
	w.Resize(fyne.NewSize(700, 400))
	w.CenterOnScreen()

	homeDir, _ := os.UserHomeDir()
	desktopDir := filepath.Join(homeDir, "Desktop")
	fmt.Println(desktopDir)

	data := binding.NewFloat()
	progressBar := widget.NewProgressBarWithData(data)
	containerProgressBar := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), progressBar, layout.NewSpacer())
	containerProgressBar.Hide()

	uploadButton := widget.NewButton("Выбрать файлы", func() {
		filesOpen := dialog.NewFolderOpen(func(dir fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			if dir != nil {
				containerProgressBar.Show()

				files, _ := dir.List()

				// debug
				fmt.Println(dir.Name())

				var filesPath []string
				for _, file := range files {
					if file.MimeType() == "image/jpeg" {
						filesPath = append(filesPath, file.Path())
					}
				}

				watermark.AddWatermarkAndMove(desktopDir, filesPath, data)

				containerProgressBar.Hide()
			}
		}, w)
		filesOpen.Show()
	})
	containerButton := container.New(layout.NewCenterLayout(), uploadButton)

	containers := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),
		container.New(layout.NewHBoxLayout(), layout.NewSpacer(), containerButton, layout.NewSpacer()),
		containerProgressBar,
		layout.NewSpacer(),
	)
	w.SetContent(containers)

	w.ShowAndRun()
}
