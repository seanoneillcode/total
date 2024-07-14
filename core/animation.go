package core

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	image        *ebiten.Image
	frames       int
	timePerFrame float64
	size         int
	// state
	frame int
	timer float64
}

func NewAnimation(image *ebiten.Image, frames int, timePerFrame float64, size int) *Animation {
	return &Animation{
		image:        image,
		frames:       frames,
		timePerFrame: timePerFrame,
		size:         size,
	}
}

func (a *Animation) Update(delta float64, game *Game) {
	a.timer = a.timer + delta
	if a.timer > a.timePerFrame {
		a.timer = a.timer - a.timePerFrame
		a.frame = a.frame + 1
		if a.frame == a.frames {
			a.frame = 0
		}
	}
}

func (a *Animation) GetImage() *ebiten.Image {
	return a.image.SubImage(image.Rect(a.frame*a.size, 0, (a.frame+1)*a.size, a.size)).(*ebiten.Image)
}
