package main

import (
	"image-cropper/handler"
	"image-cropper/repository"
	"image-cropper/service"
)

func main() {
	// 1. bikin repository (akses file & logika dasar image)
	repo := repository.NewCropRepository()

	// 2. inject repository ke service (proses bisnis: load, detect border, crop, save)
	svc := service.NewCropService(repo)

	// 3. inject service ke handler (jembatan ke layer paling atas, misalnya CLI / API)
	h := handler.NewCropHandler(svc)

	// 4. jalankan handler → otomatis trigger workflow crop image
	h.Run()
}
