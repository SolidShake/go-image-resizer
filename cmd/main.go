package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/SolidShake/go-image-resizer/internal/file"
	"github.com/SolidShake/go-image-resizer/internal/image"
)

func main() {
	a := app.New()
	w := a.NewWindow("Watermark Photo Resizer")
	w.Resize(fyne.NewSize(700, 400))
	w.CenterOnScreen()

	progressData := binding.NewFloat()
	progressBar := widget.NewProgressBarWithData(progressData)
	containerProgressBar := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), progressBar, layout.NewSpacer())
	containerProgressBar.Hide()

	uploadButton := widget.NewButton("Выбрать файлы для добавления вотермарки", func() {
		filesOpen := dialog.NewFolderOpen(func(dir fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			if dir != nil {
				containerProgressBar.Show()

				files, _ := dir.List()

				if err := image.AddWatermarkAndMove(file.GetJpegPaths(files), progressData); err != nil {
					dialog.ShowError(err, w)
				} else {
					dialog.ShowInformation("", "Фотографии успешно обработаны", w)
				}

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
