package api

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"strings"

)

func DrawText(currentImage image.Image, text string) (image.Image, error) {
	width, height := currentImage.Bounds().Max.X, currentImage.Bounds().Max.Y
	img := gg.NewContextForImage(currentImage)
	img.SetRGB(1,1,1)
	f, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}
	fSize, wrapped := calculateFontSize(f, text, img)
	face := truetype.NewFace(f, &truetype.Options{Size: fSize})
	img.SetFontFace(face)
	if wrapped {
		img.DrawStringWrapped(text, float64(width-5)/2, float64(height-5)/2, 0.5, 0.5, float64(width-10), 1, gg.AlignCenter)
	} else {
		img.DrawStringAnchored(text, float64(width-5)/2, float64(height-5)/2, 0.5, 0.5)
	}
	return img.Image(), nil
}

func calculateFontSize(f *truetype.Font, text string, img *gg.Context) (float64, bool) {
	width, height := img.Image().Bounds().Max.X, img.Image().Bounds().Max.Y
	fontSize := float64(img.Image().Bounds().Max.Y) / 3
	wrapped := false
	face := truetype.NewFace(f, &truetype.Options{Size: fontSize})
	img.SetFontFace(face)
	textWidth, _ := img.MeasureString(text)
	fSize := fontSize
	if textWidth >= float64(width-10) {
		s := (float64(width-10) / float64(textWidth)) * fontSize
		if s > 12 || !strings.Contains(text, " ") {
			fSize = s
		} else {
			wrapped = true
			words := img.WordWrap(text, float64(width - 10))
			t := ""
			for i, word := range words {
				t += word
				if i < len(words) - 1 {
					t += "\n"
				}
			}
			textWidth, textHeight := img.MeasureMultilineString(t, 1.0)
			if textHeight > textWidth && textHeight > float64(height-10) {
				fSize = (float64(height-10) / float64(textHeight)) * fontSize

			} else if textWidth > textHeight && textWidth > float64(width-10) {
				fSize = (float64(height-10) / float64(textWidth)) * fontSize
			}
		}
	}
	return fSize, wrapped
}

func ResizeImage(img image.Image, keySize int) image.Image {
	return resize.Resize(uint(keySize), uint(keySize), img, resize.Lanczos3)
}