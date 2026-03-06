package api

import (
	"image"
	"math"
	"regexp"
	"strings"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/gobolditalic"
	"golang.org/x/image/font/gofont/goitalic"
	"golang.org/x/image/font/gofont/gomedium"
	"golang.org/x/image/font/gofont/gomediumitalic"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/gofont/gomonobold"
	"golang.org/x/image/font/gofont/gomonobolditalic"
	"golang.org/x/image/font/gofont/gomonoitalic"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/gofont/gosmallcaps"
	"golang.org/x/image/font/gofont/gosmallcapsitalic"
)

const BORDER_CLEARANCE = 10

func DrawText(currentImage image.Image, text string, fontSize int, verticalAlignment, fontFace, colour string) (image.Image, error) {
	width, height := currentImage.Bounds().Max.X, currentImage.Bounds().Max.Y
	img := gg.NewContextForImage(currentImage)
	img.SetRGB(1, 1, 1)
	matched, _ := regexp.MatchString(`#?([0-9a-fA-F]{8}|[0-9a-fA-F]{6}|[0-9a-fA-F]{3})`, colour)
	if matched {
		img.SetHexColor(colour)
	}
	f, err := truetype.Parse(loadFontFace(fontFace))
	if err != nil {
		return nil, err
	}
	fSize := calculateFontSize(f, text, img)

	if fontSize != 0 {
		fSize = float64(fontSize)
	}

	face := truetype.NewFace(f, &truetype.Options{Size: fSize})
	defer face.Close()
	img.SetFontFace(face)

	lines := img.WordWrap(text, float64(width-BORDER_CLEARANCE))
	lineCount := float64(len(lines))

	if strings.Contains(text, "\n") {
		lineCount += float64(strings.Count(text, "\n") + 1)
	}

	valign, y := calculateVerticalAlignment(verticalAlignment, height, fSize, lineCount)
	img.DrawStringWrapped(text, float64(width-BORDER_CLEARANCE/2)/2, y, 0.5, valign, float64(width-BORDER_CLEARANCE), 1, gg.AlignCenter)
	return img.Image(), nil
}

// TODO Support loading fonts via fontconfig on linux and whatever the equivalent is on darwin
func loadFontFace(fontName string) []byte {
	switch fontName {
	case "bold":
		return gobold.TTF
	case "bolditalic":
		return gobolditalic.TTF
	case "italic":
		return goitalic.TTF
	case "medium":
		return gomedium.TTF
	case "mediumitalic":
		return gomediumitalic.TTF
	case "mono":
		return gomono.TTF
	case "monobold":
		return gomonobold.TTF
	case "monobolditalic":
		return gomonobolditalic.TTF
	case "monoitalic":
		return gomonoitalic.TTF
	case "smallcaps":
		return gosmallcaps.TTF
	case "smallcapsitalic":
		return gosmallcapsitalic.TTF
	case "regular":
		fallthrough
	default:
		return goregular.TTF
	}
}

func calculateVerticalAlignment(valign string, height int, fSize float64, lineCount float64) (float64, float64) {
	verticalAlignment := 0.5
	y := float64(height-BORDER_CLEARANCE/2) / 2
	if strings.ToUpper(valign) == "TOP" {
		verticalAlignment = 1.0
		y = (fSize/2)*lineCount + BORDER_CLEARANCE*lineCount
	} else if strings.ToUpper(valign) == "BOTTOM" {
		verticalAlignment = 0.0
		y = float64(height-BORDER_CLEARANCE/2) - (fSize * lineCount)
	}
	return verticalAlignment, y
}

func calculateFontSize(f *truetype.Font, text string, img *gg.Context) float64 {
	width, height := img.Image().Bounds().Max.X, img.Image().Bounds().Max.Y
	fontSize := float64(img.Image().Bounds().Max.Y) / 3
	face := truetype.NewFace(f, &truetype.Options{Size: fontSize})
	defer face.Close()
	img.SetFontFace(face)
	textWidth, _ := img.MeasureMultilineString(text, 1.0)
	fSize := fontSize
	if textWidth >= float64(width-BORDER_CLEARANCE) {
		oversizeRatio := float64(width-BORDER_CLEARANCE) / textWidth
		scaledFontSize := math.Min(oversizeRatio*fontSize, 12)
		for size := fontSize; size >= scaledFontSize; size -= 0.5 {
			if attemptFontSize(f, text, img, size, width, height) {
				return size
			}
		}
		return scaledFontSize
	}
	return fSize
}

func attemptFontSize(f *truetype.Font, text string, img *gg.Context, fSize float64, width, height int) bool {
	face := truetype.NewFace(f, &truetype.Options{Size: fSize})
	defer face.Close()
	img.SetFontFace(face)
	wrappedGroups := img.WordWrap(text, float64(width-BORDER_CLEARANCE))
	wrappedText := strings.Join(wrappedGroups, "\n")
	textWidth, textHeight := img.MeasureMultilineString(wrappedText, 1.0)
	return textHeight < float64(height-BORDER_CLEARANCE) && textWidth < float64(width-BORDER_CLEARANCE)
}

func ResizeImage(img image.Image, keySize int) image.Image {
	return resize.Resize(uint(keySize), uint(keySize), img, resize.Lanczos3)
}

func ResizeImageWH(img image.Image, width int, height int) image.Image {
	return resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
}
