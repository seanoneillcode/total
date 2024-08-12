package core

import (
	"image"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	image        *ebiten.Image
	frames       int
	timePerFrame float64
	size         int
	oneShot      bool
	// state
	frame int
	timer float64
}

func NewAnimation(image *ebiten.Image, frames int, timePerFrame float64, size int, oneShot bool) *Animation {
	return &Animation{
		image:        image,
		frames:       frames,
		timePerFrame: timePerFrame,
		size:         size,
		oneShot:      oneShot,
		frame:        rand.Intn(frames),
	}
}

func (a *Animation) Update(delta float64, game *Game) {
	if a.frame == a.frames-1 && a.oneShot {
		return
	}
	a.timer = a.timer + delta
	if a.timer > a.timePerFrame {
		a.timer = a.timer - a.timePerFrame
		a.frame = a.frame + 1
		if a.frame == a.frames {
			if !a.oneShot {
				a.frame = 0
			} else {
				a.frame = a.frames - 1
			}
		}
	}
}

func (a *Animation) GetImage() *ebiten.Image {
	return a.image.SubImage(image.Rect(a.frame*a.size, 0, (a.frame+1)*a.size, a.size)).(*ebiten.Image)
}
