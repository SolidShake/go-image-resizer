.PHONY: build run

build:
	go build cmd/main.go

run:
	go run .

static:
	fyne bundle -o bundled.go assets/watermark.png

package:
	fyne package