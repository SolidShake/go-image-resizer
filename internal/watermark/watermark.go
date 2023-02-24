package watermark

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func AddWatermarkAndMove(userDir string, files []string) {
	currentDirName := createFolderName(userDir)
	if err := os.Mkdir(currentDirName, os.ModePerm); err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if err := addWatermark(currentDirName, file); err != nil {
			fmt.Println(err)
		}
	}
}

func addWatermark(currentDirName, file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		return err
	}

	wmb, err := os.Open("watermark.png")
	if err != nil {
		return err
	}

	watermark, err := png.Decode(wmb)
	if err != nil {
		return err
	}
	defer wmb.Close()

	offset := image.Pt(200, 200)
	// missing SOI marker error if not .jgeg
	b := img.Bounds()
	m := image.NewRGBA(b)
	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	filePath := strings.Split(file, "/")
	fileName := filePath[len(filePath)-1]

	imgw, err := os.Create(currentDirName + "/" + fileName)
	if err != nil {
		return err
	}
	jpeg.Encode(imgw, m, &jpeg.Options{jpeg.DefaultQuality})
	defer imgw.Close()

	return nil
}

func createFolderName(userDir string) string {
	currentTime := time.Now().Format("2006-01-02_15-04-05")

	return filepath.Join(userDir, currentTime)
}
