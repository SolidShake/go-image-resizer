package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(700, 400))
	w.CenterOnScreen()

	homeDir, _ := os.UserHomeDir()
	desktopDir := filepath.Join(homeDir, "Desktop")
	fmt.Println(desktopDir)

	progressBar := widget.NewProgressBarInfinite()
	containerProgressBar := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), progressBar, layout.NewSpacer())
	containerProgressBar.Hide()

	uploadButton := widget.NewButton("Выбрать файлы", func() {
		filesOpen := dialog.NewFolderOpen(func(dir fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			if dir != nil {
				//
				containerProgressBar.Show()
				//
				files, _ := dir.List()
				fmt.Println(files)
				for _, file := range files {
					fmt.Println(file)
				}
				//
				time.Sleep(time.Second * 2)
				//
				containerProgressBar.Hide()
			}
		}, w)
		filesOpen.Show()
	})
	containerButton := container.New(layout.NewCenterLayout(), uploadButton)

	containers := container.New(
		layout.NewVBoxLayout(),
		container.New(layout.NewHBoxLayout(), layout.NewSpacer(), containerButton, layout.NewSpacer()),
		containerProgressBar,
	)
	w.SetContent(containers)

	w.ShowAndRun()
}
