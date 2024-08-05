package core

import (
	"total/common"

	"github.com/hajimehoshi/ebiten/v2"
)

type Cursor struct {
	x      int
	y      int
	images map[string]*Animation
	state  string
}

func NewCursor(game *Game) *Cursor {
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	cx, cy := ebiten.CursorPosition()
	return &Cursor{
		x: cx,
		y: cy,
		images: map[string]*Animation{
			"cursor":      NewAnimation(game.images["cursor"], 1, 1, 16),
			"cursor-move": NewAnimation(game.images["cursor-move"], 4, 0.1, 16),
		},
		state: "cursor",
	}
}

func (r *Cursor) Update(delta float64, game *Game) {
	cx, cy := ebiten.CursorPosition()
	r.x = cx
	r.y = cy
	r.state = "cursor"
	if game.selectedUnit != nil {
		r.state = "cursor-move"
	}
	r.images[r.state].Update(delta, game)
}

func (r *Cursor) Draw(camera *Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.x/common.Scale)-8, float64(r.y/common.Scale)-8)
	op.GeoM.Scale(common.Scale, common.Scale)
	camera.DrawUI(r.images[r.state].GetImage(), op, forgroundLayer)
}
