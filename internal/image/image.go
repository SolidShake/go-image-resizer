package image

import (
	"fmt"

	"fyne.io/fyne/v2/data/binding"
	"github.com/SolidShake/go-image-resizer/internal/file"
	"github.com/SolidShake/go-image-resizer/internal/watermark"
)

func AddWatermarkAndMove(userDir string, files []string, progressData binding.Float) {
	newFolder, err := file.CreateFolder(userDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	total := 0.0
	part := 1.0 / float64(len(files))

	for _, file := range files {
		fmt.Println(file)
		if err := watermark.AddWatermark(newFolder, file); err != nil {
			fmt.Println(file, err)
		}

		total += part
		progressData.Set(total)
	}
}
