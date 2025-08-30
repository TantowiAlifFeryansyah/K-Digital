package repository

import (
	"image"
	"os"

	_ "image/jpeg"
	"image/png"

	"image-cropper/entity"
	"image-cropper/helper"
)

// load image, save image, dan detect border
type CropRepository interface {
	LoadImage(path string) (image.Image, error)
	SaveImage(img image.Image, path string) error
	// cari kotak border hitam + list koordinat pikselnya
	DetectBorder(img image.Image) (*entity.BorderBox, []image.Point, error)
}

// struct kosong yang jadi implementasi dari interface di atas
type cropRepository struct{}

// buat instance baru dari CropRepository
func NewCropRepository() CropRepository {
	return &cropRepository{}
}

// buka file dari path, decode jadi image.Image
func (r *cropRepository) LoadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// SaveImage: simpan image ke file PNG
func (r *cropRepository) SaveImage(img image.Image, path string) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()
	return png.Encode(out, img)
}

// DetectBorder:
//  1. Cari seed → piksel hitam pertama (scan dari kiri-atas).
//  2. Dari seed, lakukan BFS (8 arah) buat kumpulin semua piksel border yang nyambung.
//  3. Catat minX, minY, maxX, maxY buat dapetin bounding box.
func (r *cropRepository) DetectBorder(img image.Image) (*entity.BorderBox, []image.Point, error) {
	b := img.Bounds()
	w := b.Dx()
	h := b.Dy()

	// cari seed piksel hitam pertama
	found := false
	seed := image.Point{}
	for y := b.Min.Y; y < b.Max.Y && !found; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			if helper.IsBlack(img.At(x, y)) {
				seed = image.Point{X: x, Y: y}
				found = true
				break
			}
		}
	}
	if !found {
		// jika border tidak ada
		return nil, nil, os.ErrNotExist
	}

	// array visited untuk penanda piksel yang udah dicek
	visited := make([]bool, w*h)
	idx := func(x, y int) int { return (y-b.Min.Y)*w + (x - b.Min.X) }

	// queue BFS, mulai dari seed
	queue := make([]image.Point, 0, 1024)
	queue = append(queue, seed)
	visited[idx(seed.X, seed.Y)] = true

	// inisialisasi bounding box
	minX, minY := seed.X, seed.Y
	maxX, maxY := seed.X, seed.Y
	points := make([]image.Point, 0, 8192)
	points = append(points, seed)

	// arah gerakan BFS (8 tetangga: kiri, kanan, atas, bawah, + diagonal)
	neighbors := [8][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}

	// proses BFS
	for qi := 0; qi < len(queue); qi++ {
		p := queue[qi]
		for _, d := range neighbors {
			nx, ny := p.X+d[0], p.Y+d[1]
			// skip jika di luar bounds
			if nx < b.Min.X || nx >= b.Max.X || ny < b.Min.Y || ny >= b.Max.Y {
				continue
			}
			id := idx(nx, ny)
			if visited[id] {
				continue
			}
			// hanya proses kalau pikselnya hitam
			if !helper.IsBlack(img.At(nx, ny)) {
				continue
			}
			visited[id] = true
			qp := image.Point{X: nx, Y: ny}
			queue = append(queue, qp)
			points = append(points, qp)

			// update bounding box
			if nx < minX {
				minX = nx
			}
			if ny < minY {
				minY = ny
			}
			if nx > maxX {
				maxX = nx
			}
			if ny > maxY {
				maxY = ny
			}
		}
	}

	// kalau jumlah piksel kosong, dianggap bukan border valid
	if len(points) == 0 {
		return nil, nil, os.ErrNotExist
	}

	// return bounding box + list piksel border
	box := &entity.BorderBox{
		MinX: minX,
		MinY: minY,
		MaxX: maxX,
		MaxY: maxY,
	}
	return box, points, nil
}
