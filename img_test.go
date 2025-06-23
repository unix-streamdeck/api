package api

import (
    "image"
    "testing"
)

func TestResizeImage(t *testing.T) {
    img := image.NewRGBA(image.Rect(0, 0, 1000, 1000)).SubImage(image.Rect(0, 0, 1000, 1000))
    size := 72
    img = ResizeImage(img, size)
    maxBounds := img.Bounds().Max
    if maxBounds.X != size || maxBounds.Y != size {
        t.Errorf("Size should be 72 px but got %dx%d", maxBounds.X, maxBounds.Y)
        t.FailNow()
    }
}

func TestResizeImageWH(t *testing.T) {
    img := image.NewRGBA(image.Rect(0, 0, 1000, 1000)).SubImage(image.Rect(0, 0, 1000, 1000))
    width := 100
    height := 50
    img = ResizeImageWH(img, width, height)
    maxBounds := img.Bounds().Max
    if maxBounds.X != width || maxBounds.Y != height {
        t.Errorf("Size should be %dx%d but got %dx%d", width, height, maxBounds.X, maxBounds.Y)
        t.FailNow()
    }
}

func TestDrawText(t *testing.T) {
    // Create a test image
    img := image.NewRGBA(image.Rect(0, 0, 100, 100))

    // Test with default font size and center alignment
    text := "Test"
    fontSize := 0 // Use default font size
    fontAlignment := "center"

    result, err := DrawText(img, text, fontSize, fontAlignment)
    if err != nil {
        t.Errorf("DrawText returned an error: %v", err)
        t.FailNow()
    }

    // Check that the result is not nil and has the same dimensions as the input
    if result == nil {
        t.Error("DrawText returned nil image")
        t.FailNow()
    }

    maxBounds := result.Bounds().Max
    if maxBounds.X != img.Bounds().Max.X || maxBounds.Y != img.Bounds().Max.Y {
        t.Errorf("Result image size should be %dx%d but got %dx%d", 
            img.Bounds().Max.X, img.Bounds().Max.Y, maxBounds.X, maxBounds.Y)
        t.FailNow()
    }

    // Test with specific font size
    fontSize = 20
    result, err = DrawText(img, text, fontSize, fontAlignment)
    if err != nil {
        t.Errorf("DrawText returned an error: %v", err)
        t.FailNow()
    }

    // Test with top alignment
    fontAlignment = "top"
    result, err = DrawText(img, text, fontSize, fontAlignment)
    if err != nil {
        t.Errorf("DrawText returned an error: %v", err)
        t.FailNow()
    }

    // Test with bottom alignment
    fontAlignment = "bottom"
    result, err = DrawText(img, text, fontSize, fontAlignment)
    if err != nil {
        t.Errorf("DrawText returned an error: %v", err)
        t.FailNow()
    }

    // Test with multiline text
    text = "Line 1\nLine 2"
    result, err = DrawText(img, text, fontSize, fontAlignment)
    if err != nil {
        t.Errorf("DrawText returned an error: %v", err)
        t.FailNow()
    }
}
