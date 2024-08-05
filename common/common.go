package common

import (
	"bytes"
	"errors"
	"image"
	"io/ioutil"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	_ "image/png" // evil, required for decoder to 'know' what a png is...
)

const (
	ScreenWidth  = 360
	ScreenHeight = 240
	Scale        = 4
	TextScale    = 2
)

var NormalEscapeError = errors.New("normal escape termination")

func LoadImage(imageFileName string) *ebiten.Image {
	return loadImage("res/" + imageFileName)
}

func loadImage(imageFileName string) *ebiten.Image {
	b, err := ioutil.ReadFile(imageFileName)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}

func Overlap(x1, y1, s1, x2, y2, s2 float64) bool {
	if x2 > x1+s1 || x2+s2 < x1 {
		return false
	}
	if y2 > y1+s1 || y2+s2 < y1 {
		return false
	}
	return true
}

func Normalize(x float64, y float64) (float64, float64) {
	var nx, ny float64
	norm2 := x*x + y*y
	if norm2 == 0 {
		return nx, ny
	}
	ratio := (1 / math.Sqrt(norm2))
	nx = x * ratio
	ny = y * ratio
	return nx, ny
}
