package helper

import (
	"fmt"
	"image"
	"image/color"
	"os"
)

// untuk simpan koordinat piksel ke file log
// format tiap baris: x=..., y=...
func SaveLog(points []image.Point, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, p := range points {
		_, err := fmt.Fprintf(file, "x=%d, y=%d\n", p.X, p.Y)
		if err != nil {
			return err
		}
	}
	return nil
}

// cek apakah warna piksel pure black (0,0,0)
func IsBlack(c color.Color) bool {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	return rgba.R == 0 && rgba.G == 0 && rgba.B == 0
}
