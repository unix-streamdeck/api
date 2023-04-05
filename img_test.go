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