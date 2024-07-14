package core

import (
	"total/common"

	"github.com/hajimehoshi/ebiten/v2"
)

const HalfScreenWidth = common.ScreenWidth / 2
const HalfScreenHeight = common.ScreenHeight / 2

type Camera struct {
	x      float64
	y      float64
	screen *ebiten.Image
}

func (r *Camera) DrawImage(image *ebiten.Image, op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(-r.x+HalfScreenWidth, -r.y+HalfScreenHeight)
	op.GeoM.Scale(common.Scale, common.Scale)
	r.screen.DrawImage(image, op)
}
