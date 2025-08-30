package service

import (
	"fmt"
	"image"
	"image-cropper/helper"
	"image-cropper/repository"
)

// fitur utama untuk crop image
type CropService interface {
	CropImage(inputPath, outputPath, logPath string) error
}

// penghubung antara repository dengan layer yang manggil
type cropService struct {
	repo repository.CropRepository
}

// instance baru dari service, sekaligus inject repository yang dipakai
func NewCropService(repo repository.CropRepository) CropService {
	return &cropService{repo: repo}
}

// workflow utama proses crop gambar
func (s *cropService) CropImage(inputPath, outputPath, logPath string) error {
	// 1. load image dari file
	img, err := s.repo.LoadImage(inputPath)
	if err != nil {
		return err
	}

	// 2. deteksi border hitam → dappatkan bounding box & list koordinat piksel
	bbox, points, err := s.repo.DetectBorder(img)
	if err != nil {
		return err
	}

	// 3. print info bounding box ke console (debug)
	fmt.Printf("Border detected: MinX=%d, MinY=%d, MaxX=%d, MaxY=%d\n",
		bbox.MinX, bbox.MinY, bbox.MaxX, bbox.MaxY)

	// 4. tentukan rectangle crop sesuai bounding box
	// gunakan bbox.MaxX+1 & bbox.MaxY+1 karena koordinat Rect eksklusif di Go
	rect := image.Rect(bbox.MinX, bbox.MinY, bbox.MaxX+1, bbox.MaxY+1)

	// 5. lakukan crop → ambil sub-image sesuai rect
	cropped := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(rect)

	// 6. simpan hasil crop ke file output
	if err := s.repo.SaveImage(cropped, outputPath); err != nil {
		return err
	}

	// 7. simpan log koordinat border ke file (opsional)
	if logPath != "" {
		if err := helper.SaveLog(points, logPath); err != nil {
			return err
		}
	}

	fmt.Printf("✅ Cropped image saved at %s\n", outputPath)
	return nil
}
