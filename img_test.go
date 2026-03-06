package api

import (
	"image"
	"image/draw"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResizeImage(t *testing.T) {
	assertions := assert.New(t)
	width := rand.Intn(255) + 1
	height := rand.Intn(255) + 1
	iconSize := rand.Intn(255) + 1
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), image.Black, image.ZP, draw.Src)
	resizedImage := ResizeImage(img, iconSize)
	assertions.Equal(resizedImage.Bounds().Max.X, iconSize)
	assertions.Equal(resizedImage.Bounds().Max.Y, iconSize)
}

func TestResizeImageWH(t *testing.T) {
	assertions := assert.New(t)
	width := rand.Intn(255) + 1
	height := rand.Intn(255) + 1
	newWidth := rand.Intn(255) + 1
	newHeight := rand.Intn(255) + 1
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), image.Black, image.ZP, draw.Src)
	resizedImage := ResizeImageWH(img, newWidth, newHeight)
	assertions.Equal(resizedImage.Bounds().Max.X, newWidth)
	assertions.Equal(resizedImage.Bounds().Max.Y, newHeight)
}
