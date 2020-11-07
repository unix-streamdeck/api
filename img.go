package api

import (
	"errors"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"strings"
)

func DrawText(currentImage image.Image, text string, fontSize int, fontAlignment string) (image.Image, error) {
	width, height := currentImage.Bounds().Max.X, currentImage.Bounds().Max.Y
	img := gg.NewContextForImage(currentImage)
	img.SetRGB(1,1,1)
	f, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}
	fSize, wrapped := calculateFontSize(f, text, img)
	if fontSize != 0 {
		fSize = float64(fontSize)
	}
	face := truetype.NewFace(f, &truetype.Options{Size: fSize})
	img.SetFontFace(face)
	if fontAlignment == "" {
		if wrapped {
			img.DrawStringWrapped(text, float64(width-5)/2, float64(height-5)/2, 0.5, 0.5, float64(width-10), 1, gg.AlignCenter)
		} else {
			img.DrawStringAnchored(text, float64(width-5)/2, float64(height-5)/2, 0.5, 0.5)
		}
	} else {
		verticalAlignment := 0.5
		y := float64(height-5)/2
		if strings.ToUpper(fontAlignment) == "TOP" {
			verticalAlignment = 0.0
			y = (fSize / 2) + 10
		} else if strings.ToUpper(fontAlignment) == "BOTTOM" {
			verticalAlignment = 1.0
			y = float64(height - 5) - fSize
		} else if strings.ToUpper(fontAlignment) != "MIDDLE" {
			return nil, errors.New("Invalid vertical alignment")
		}

		img.DrawStringAnchored(text, float64(width-5)/2, y, 0.5, verticalAlignment)
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