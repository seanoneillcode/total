package core

import "github.com/hajimehoshi/ebiten/v2"

type Decor struct {
	x         float64
	y         float64
	z         int
	animation *Animation
}

func (r *Decor) Draw(camera *Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(r.x, r.y)
	camera.DrawImage(r.animation.GetImage(), op, backgroundLayer)
}
