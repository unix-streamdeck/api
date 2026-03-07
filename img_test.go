package api

import (
	"image"
	"image/draw"
	"math/rand"
	"testing"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/stretchr/testify/assert"
	"github.com/unix-streamdeck/api/v2/mocks/mock_api"
	"go.uber.org/mock/gomock"
	"golang.org/x/image/font/gofont/goregular"
)

func TestDrawText(t *testing.T) {
	assertions := assert.New(t)
	ctrl := gomock.NewController(t)

	context := mock_api.NewMockIContext(ctrl)

	context.EXPECT().SetRGB(1.0, 1.0, 1.0).Times(1)

	context.EXPECT().Width().Return(72).Times(2)
	context.EXPECT().Height().Return(72).Times(1)

	context.EXPECT().SetFontFace(gomock.Any()).Times(2)

	context.EXPECT().MeasureMultilineString("Test", 1.0).Return(20.0, 24.0).Times(1)

	context.EXPECT().WordWrap("Test", 62.0).Return([]string{"Test"}).Times(1)

	context.EXPECT().DrawStringWrapped("Test", 36.0, 36.0, 0.5, 0.5, 62.0, 1.0, gg.AlignCenter)

	mockImg := setupImage(72, 72)

	context.EXPECT().Image().Return(mockImg).Times(1)

	img, err := drawText(context, "Test", DrawTextOptions{})

	assertions.Equal(mockImg, img)

	assertions.Nil(err)

}

func Test_calculateVerticalAlignment_Center(t *testing.T) {
	assertions := assert.New(t)
	alignment, anchor := calculateVerticalAlignment(Center, 80)
	assertions.Equal(0.5, alignment)
	assertions.Equal(40.0, anchor)
}

func Test_calculateVerticalAlignment_Top(t *testing.T) {
	assertions := assert.New(t)
	alignment, anchor := calculateVerticalAlignment(Top, 80)
	assertions.Equal(0.0, alignment)
	assertions.Equal(5.0, anchor)
}

func Test_calculateVerticalAlignment_Bottom(t *testing.T) {
	assertions := assert.New(t)
	alignment, anchor := calculateVerticalAlignment(Bottom, 80)
	assertions.Equal(1.0, alignment)
	assertions.Equal(75.0, anchor)
}

func Test_calculateFontSize_SingleWord(t *testing.T) {
	assertions := assert.New(t)
	img := setupImage(72, 72)
	ggImg := gg.NewContextForImage(img)
	ggImg.SetHexColor("#FFF")
	f, _ := truetype.Parse(goregular.TTF)
	assertions.Equal(24.0, calculateFontSize(f, "Test", ggImg))
}

func Test_calculateFontSize_MultiLine(t *testing.T) {
	assertions := assert.New(t)
	img := setupImage(72, 72)
	ggImg := gg.NewContextForImage(img)
	ggImg.SetHexColor("#FFF")
	f, _ := truetype.Parse(goregular.TTF)
	assertions.Equal(24.0, calculateFontSize(f, "Lines Test", ggImg))
}

func Test_calculateFontSize_LongMultiLine(t *testing.T) {
	assertions := assert.New(t)
	img := setupImage(72, 72)
	ggImg := gg.NewContextForImage(img)
	ggImg.SetHexColor("#FFF")
	f, _ := truetype.Parse(goregular.TTF)
	assertions.Equal(15.5, calculateFontSize(f, "Multiline Overflow", ggImg))
}

func Test_attemptFontSize_SingleLine(t *testing.T) {
	assertions := assert.New(t)
	img := setupImage(72, 72)
	ggImg := gg.NewContextForImage(img)
	ggImg.SetHexColor("#FFF")
	f, _ := truetype.Parse(goregular.TTF)
	assertions.True(attemptFontSize(f, "Test", ggImg, 24.0))
}

func Test_attemptFontSize_MultiLine(t *testing.T) {
	assertions := assert.New(t)
	img := setupImage(72, 72)
	ggImg := gg.NewContextForImage(img)
	ggImg.SetHexColor("#FFF")
	f, _ := truetype.Parse(goregular.TTF)
	assertions.True(attemptFontSize(f, "Lines Test", ggImg, 24.0))
}

func Test_attemptFontSize_Overflow(t *testing.T) {
	assertions := assert.New(t)
	img := setupImage(72, 72)
	ggImg := gg.NewContextForImage(img)
	ggImg.SetHexColor("#FFF")
	f, _ := truetype.Parse(goregular.TTF)
	assertions.False(attemptFontSize(f, "Muiltiline Overflow", ggImg, 24.0))
}

func TestResizeImage(t *testing.T) {
	assertions := assert.New(t)
	width := rand.Intn(255) + 1
	height := rand.Intn(255) + 1
	iconSize := rand.Intn(255) + 1
	img := setupImage(width, height)
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
	img := setupImage(width, height)
	resizedImage := ResizeImageWH(img, newWidth, newHeight)
	assertions.Equal(resizedImage.Bounds().Max.X, newWidth)
	assertions.Equal(resizedImage.Bounds().Max.Y, newHeight)
}

func setupImage(width int, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), image.Black, image.ZP, draw.Src)
	return img
}
