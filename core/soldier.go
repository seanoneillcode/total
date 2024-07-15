package core

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Soldier struct {
	x          float64
	y          float64
	tx         float64
	ty         float64
	speed      float64
	isFlip     bool
	animations map[string]*Animation
	state      string
}

func NewSoldier(game *Game, x float64, y float64) *Soldier {
	return &Soldier{
		animations: map[string]*Animation{
			"idle": NewAnimation(game.images["soldier-idle"], 4, 0.2, 16),
			"walk": NewAnimation(game.images["soldier-walk"], 4, 0.1, 16),
		},
		x:     x,
		y:     y,
		tx:    x,
		ty:    y,
		speed: 30,
		state: "idle",
	}
}

func (r *Soldier) Update(delta float64, game *Game) {
	isMoving := false
	if math.Abs(r.tx-r.x) < (2 * r.speed * delta) {
		r.x = r.tx
	}
	if math.Abs(r.ty-r.y) < (2 * r.speed * delta) {
		r.y = r.ty
	}
	if r.x < r.tx {
		r.x = r.x + r.speed*delta
		isMoving = true
		r.isFlip = false
	}
	if r.x > r.tx {
		r.x = r.x - r.speed*delta
		isMoving = true
		r.isFlip = true
	}
	if r.y < r.ty {
		r.y = r.y + r.speed*delta
		isMoving = true
	}
	if r.y > r.ty {
		r.y = r.y - r.speed*delta
		isMoving = true
	}
	r.state = "idle"
	if isMoving {
		r.state = "walk"
	}
	r.animations[r.state].Update(delta, game)

}

func (r *Soldier) Draw(camera *Camera) {
	op := &ebiten.DrawImageOptions{}
	if r.isFlip {
		op.GeoM.Scale(-1, 1)
	}
	op.GeoM.Translate(r.x, r.y)
	if r.isFlip {
		op.GeoM.Translate(16, 0)
	}
	animation := r.animations[r.state]
	camera.DrawImage(animation.GetImage(), op)
}
