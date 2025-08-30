package handler

import (
	"fmt"
	"image-cropper/constants"
	"image-cropper/response"
	"image-cropper/service"
)

// penghubung antara service dengan layer yang manggil
type CropHandler struct {
	service service.CropService
}

// instance baru dari handler, sekaligus inject service yang dipakai
func NewCropHandler(service service.CropService) *CropHandler {
	return &CropHandler{service: service}
}

// method utama
func (h *CropHandler) Run() {
	// file input yang mau di-crop
	input := "uploads/image.png"
	// hasil output setelah dicrop
	output := "output.png"
	// file log untuk simpan info border
	log := "border.log"

	// panggil service crop image
	err := h.service.CropImage(input, output, log)
	if err != nil {
		// response jika gagal
		fmt.Println(response.NewResponse(constants.InternalErrorCode, err.Error(), nil))
		return
	}

	// response akhir + log
	fmt.Println(response.NewResponse(constants.SuccessCode, constants.ImagesCroppedSuccess, map[string]string{
		"output": output,
		"log":    log,
	}))
}
