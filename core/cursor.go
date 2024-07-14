package core

import (
	"total/common"

	"github.com/hajimehoshi/ebiten/v2"
)

type Cursor struct {
	x     int
	y     int
	image *ebiten.Image
}

func NewCursor(game *Game) *Cursor {
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	cx, cy := ebiten.CursorPosition()
	return &Cursor{
		x:     cx,
		y:     cy,
		image: game.images["cursor"],
	}
}

func (r *Cursor) Update() {
	cx, cy := ebiten.CursorPosition()
	r.x = cx
	r.y = cy
}

func (r *Cursor) Draw(camera *Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.x/common.Scale)-8, float64(r.y/common.Scale)-8)
	op.GeoM.Scale(common.Scale, common.Scale)
	camera.screen.DrawImage(r.image, op)
}
