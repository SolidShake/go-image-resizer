.PHONY: build run

build:
	go build cmd/main.go

run:
	go run cmd/main.go

release-macos:
	fyne release -os ios -certificate "Apple Distribution" -profile "My App Distribution" -appID "com.example.myapp"