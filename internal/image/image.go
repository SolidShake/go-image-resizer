package image

import (
	"errors"
	"fmt"

	"fyne.io/fyne/v2/data/binding"
	"github.com/SolidShake/go-image-resizer/internal/file"
	"github.com/SolidShake/go-image-resizer/internal/watermark"
)

func AddWatermarkAndMove(files []string, progressData binding.Float) error {
	newFolder, err := file.CreateNewFolder()
	if err != nil {
		return err
	}

	total := 0.0
	part := 1.0 / float64(len(files))

	var watermarkError error

	for _, file := range files {
		// debug
		fmt.Println(file)
		if err := watermark.AddWatermark(newFolder, file); err != nil {
			watermarkError = errors.Join(watermarkError, fmt.Errorf("file: %s, error: %w", file, err))
		}

		total += part
		progressData.Set(total)
	}

	return watermarkError
}
