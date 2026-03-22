package api

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"math/rand"
	"testing"

	"github.com/golang/freetype/truetype"
	"github.com/stretchr/testify/assert"
	"github.com/unix-streamdeck/api/v2/mocks/mock_api"
	"github.com/unix-streamdeck/gg"
	"go.uber.org/mock/gomock"
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

func TestDrawText_WithColor(t *testing.T) {
	assertions := assert.New(t)
	ctrl := gomock.NewController(t)

	context := mock_api.NewMockIContext(ctrl)

	context.EXPECT().SetRGB(1.0, 1.0, 1.0).Times(1)
	context.EXPECT().SetHexColor("#FF0000").Times(1)

	context.EXPECT().Width().Return(72).Times(2)
	context.EXPECT().Height().Return(72).Times(1)

	context.EXPECT().SetFontFace(gomock.Any()).Times(2)

	context.EXPECT().MeasureMultilineString("Test", 1.0).Return(20.0, 24.0).Times(1)

	context.EXPECT().WordWrap("Test", 62.0).Return([]string{"Test"}).Times(1)

	context.EXPECT().DrawStringWrapped("Test", 36.0, 36.0, 0.5, 0.5, 62.0, 1.0, gg.AlignCenter)

	mockImg := setupImage(72, 72)

	context.EXPECT().Image().Return(mockImg).Times(1)

	img, err := drawText(context, "Test", DrawTextOptions{Colour: "#FF0000"})

	assertions.Equal(mockImg, img)
	assertions.Nil(err)
}

func TestDrawText_WithFontSize(t *testing.T) {
	assertions := assert.New(t)
	ctrl := gomock.NewController(t)

	context := mock_api.NewMockIContext(ctrl)

	context.EXPECT().SetRGB(1.0, 1.0, 1.0).Times(1)

	context.EXPECT().Width().Return(72).Times(2)
	context.EXPECT().Height().Return(72).Times(1)

	context.EXPECT().SetFontFace(gomock.Any()).Times(2)

	context.EXPECT().MeasureMultilineString("Test", 1.0).Return(25.0, 12.0).Times(1)

	context.EXPECT().WordWrap("Test", 62.0).Return([]string{"Test"}).Times(1)

	context.EXPECT().DrawStringWrapped("Test", 36.0, 36.0, 0.5, 0.5, 62.0, 1.0, gg.AlignCenter)

	mockImg := setupImage(72, 72)

	context.EXPECT().Image().Return(mockImg).Times(1)

	img, err := drawText(context, "Test", DrawTextOptions{FontSize: 16})

	assertions.Equal(mockImg, img)
	assertions.Nil(err)
}

func TestDrawText_WithNewlines(t *testing.T) {
	assertions := assert.New(t)
	ctrl := gomock.NewController(t)

	context := mock_api.NewMockIContext(ctrl)

	context.EXPECT().SetRGB(1.0, 1.0, 1.0).Times(1)

	context.EXPECT().Width().Return(72).Times(2)
	context.EXPECT().Height().Return(72).Times(1)

	context.EXPECT().SetFontFace(gomock.Any()).Times(2)

	context.EXPECT().MeasureMultilineString("Line1\nLine2", 1.0).Return(30.0, 48.0).Times(1)

	context.EXPECT().WordWrap("Line1\nLine2", 62.0).Return([]string{"Line1", "Line2"}).Times(1)

	context.EXPECT().DrawStringWrapped("Line1\nLine2", 36.0, 36.0, 0.5, 0.5, 62.0, 1.0, gg.AlignCenter)

	mockImg := setupImage(72, 72)

	context.EXPECT().Image().Return(mockImg).Times(1)

	img, err := drawText(context, "Line1\nLine2", DrawTextOptions{})

	assertions.Equal(mockImg, img)
	assertions.Nil(err)
}

func Test_loadFontFace_Bold(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(gobold.TTF, loadFontFace("bold"))
}

func Test_loadFontFace_BoldItalic(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(gobolditalic.TTF, loadFontFace("bolditalic"))
}

func Test_loadFontFace_Italic(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(goitalic.TTF, loadFontFace("italic"))
}

func Test_loadFontFace_Medium(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(gomedium.TTF, loadFontFace("medium"))
}

func Test_loadFontFace_MediumItalic(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(gomediumitalic.TTF, loadFontFace("mediumitalic"))
}

func Test_loadFontFace_Mono(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(gomono.TTF, loadFontFace("mono"))
}

func Test_loadFontFace_MonoBold(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(gomonobold.TTF, loadFontFace("monobold"))
}

func Test_loadFontFace_MonoBoldItalic(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(gomonobolditalic.TTF, loadFontFace("monobolditalic"))
}

func Test_loadFontFace_MonoItalic(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(gomonoitalic.TTF, loadFontFace("monoitalic"))
}

func Test_loadFontFace_SmallCaps(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(gosmallcaps.TTF, loadFontFace("smallcaps"))
}

func Test_loadFontFace_SmallCapsItalic(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(gosmallcapsitalic.TTF, loadFontFace("smallcapsitalic"))
}

func Test_loadFontFace_Regular(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(goregular.TTF, loadFontFace("regular"))
}

func Test_loadFontFace_Default(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(goregular.TTF, loadFontFace("unknown"))
}

func TestHexColor(t *testing.T) {
	assertions := assert.New(t)
	result := HexColor("#FF0000")
	expected := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	assertions.Equal(expected, result)
}

func TestHexColor_Green(t *testing.T) {
	assertions := assert.New(t)
	result := HexColor("#00FF00")
	expected := color.RGBA{R: 0, G: 255, B: 0, A: 255}
	assertions.Equal(expected, result)
}

func TestHexColor_Blue(t *testing.T) {
	assertions := assert.New(t)
	result := HexColor("#0000FF")
	expected := color.RGBA{R: 0, G: 0, B: 255, A: 255}
	assertions.Equal(expected, result)
}

func TestLayerImages_Success(t *testing.T) {
	assertions := assert.New(t)
	img1 := setupImage(72, 72)
	draw.Draw(img1, img1.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 255}}, image.Point{}, draw.Src)

	img2 := image.NewRGBA(image.Rect(0, 0, 72, 72))
	draw.Draw(img2, img2.Bounds(), &image.Uniform{color.RGBA{0, 255, 0, 128}}, image.Point{}, draw.Src)

	result, err := LayerImages(72, 72, img1, img2)

	log.Println()

	assertions.NotNil(result)
	assertions.Nil(err)
	assertions.Equal(72, result.Bounds().Dx())
	assertions.Equal(72, result.Bounds().Dy())

	assertions.Equal(color.RGBA{127, 255, 0, 255}, result.At(0, 0))

	rgba, ok := result.(*image.RGBA)
	assertions.True(ok)
	assertions.NotNil(rgba)
}

func TestLayerImages_NoImages(t *testing.T) {
	assertions := assert.New(t)

	result, err := LayerImages(72, 72)

	assertions.Nil(result)
	assertions.NotNil(err)
	assertions.Equal("no images supplied", err.Error())
}

func TestLayerImages_NilImages(t *testing.T) {
	assertions := assert.New(t)

	result, err := LayerImages(72, 72, nil, nil)

	assertions.Nil(result)
	assertions.NotNil(err)
	assertions.Equal("no valid images supplied", err.Error())
}

func TestLayerImages_WrongSize(t *testing.T) {
	assertions := assert.New(t)
	img1 := setupImage(50, 50)
	img2 := setupImage(60, 60)

	result, err := LayerImages(72, 72, img1, img2)

	assertions.Nil(result)
	assertions.NotNil(err)
	assertions.Equal("no valid images supplied", err.Error())
}

func TestLayerImages_MixedValidInvalid(t *testing.T) {
	assertions := assert.New(t)
	img1 := setupImage(72, 72)
	img2 := setupImage(50, 50)

	result, err := LayerImages(72, 72, img1, img2)

	assertions.NotNil(result)
	assertions.Nil(err)
}

func TestSubImage(t *testing.T) {
	assertions := assert.New(t)
	img := setupImage(100, 100)

	result := SubImage(img, 10, 10, 50, 50)

	assertions.NotNil(result)
	assertions.Equal(40, result.Bounds().Dx())
	assertions.Equal(40, result.Bounds().Dy())
	assertions.Equal(10, result.Bounds().Min.X)
	assertions.Equal(10, result.Bounds().Min.Y)
	assertions.Equal(50, result.Bounds().Max.X)
	assertions.Equal(50, result.Bounds().Max.Y)
}
func setupImage(width int, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), image.Black, image.ZP, draw.Src)
	return img
}
