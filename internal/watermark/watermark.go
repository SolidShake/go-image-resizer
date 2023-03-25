package watermark

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func AddWatermark(currentDirName, file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		return err
	}

	wmb, err := os.Open("assets/watermark.png")
	if err != nil {
		return err
	}

	watermark, err := png.Decode(wmb)
	if err != nil {
		return err
	}
	defer wmb.Close()

	x := img.Bounds().Dx()/2 - watermark.Bounds().Dx()/2
	y := img.Bounds().Dy() / 2

	offset := image.Pt(x, y)
	// missing SOI marker error if not .jpeg
	b := img.Bounds()
	m := image.NewRGBA(b)
	draw.Draw(m, b, img, image.Point{}, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.Point{}, draw.Over)

	filePath := strings.Split(file, "/")
	fileName := filePath[len(filePath)-1]

	imgw, err := os.Create(currentDirName + "/" + fileName)
	if err != nil {
		return err
	}
	jpeg.Encode(imgw, m, &jpeg.Options{Quality: jpeg.DefaultQuality})
	defer imgw.Close()

	return nil
}
