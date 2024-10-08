package core

import (
	"math"
	"math/rand"
	"total/common"

	"github.com/hajimehoshi/ebiten/v2"
)

type Soldier struct {
	x              float64
	y              float64
	tx             float64
	ty             float64
	offsetx        float64
	offsety        float64
	size           float64
	speed          float64
	isFlip         bool
	animations     map[string]*Animation
	state          string
	isSelected     bool
	selection      *ebiten.Image
	shadow         *ebiten.Image
	unitRes        UnitResource
	unitStats      UnitStats
	selectionTimer float64
}

func NewSoldier(game *Game, x float64, y float64, soldierType string) *Soldier {
	stats := game.stats.GetUnitStats(soldierType)
	unitRes := game.resources.GetUnitResource(soldierType)
	return &Soldier{
		animations: map[string]*Animation{
			"idle": NewAnimation(game.resources.GetImage(unitRes.Idle), 4, 0.1, unitRes.Size, false),
			"walk": NewAnimation(game.resources.GetImage(unitRes.Walk), 4, 0.1, unitRes.Size, false),
		},
		selection: game.resources.GetImage("selection"),
		shadow:    game.resources.GetImage("unit-shadow"),
		x:         x,
		y:         y,
		tx:        x,
		ty:        y,
		offsetx:   float64(unitRes.Size / 2),
		offsety:   float64(unitRes.Size - 2),
		size:      float64(unitRes.Size),
		speed:     stats.Speed,
		state:     "idle",
		unitRes:   unitRes,
		unitStats: stats,
	}
}

func (r *Soldier) Update(delta float64, game *Game) {
	if r.state == "die" {
		return
	}
	if r.selectionTimer > 0 {
		r.selectionTimer = r.selectionTimer - delta
	}
	isMoving := false
	if math.Abs(r.tx-r.x) < (2 * r.speed) {
		r.x = r.tx
	}
	if math.Abs(r.ty-r.y) < (2 * r.speed) {
		r.y = r.ty
	}

	var dirx float64
	var diry float64
	if r.x != r.tx || r.y != r.ty {
		dirx = r.tx - r.x
		diry = r.ty - r.y
		isMoving = true

		nx, ny := common.Normalize(dirx, diry)
		r.x = r.x + (r.speed * nx)
		r.y = r.y + (r.speed * ny)
	}

	if r.x < r.tx {
		r.isFlip = false
	}

	if r.x > r.tx {
		r.isFlip = true
	}
	r.state = "idle"
	if isMoving {
		r.state = "walk"
	}
	r.animations[r.state].Update(delta, game)
}

func (r *Soldier) Draw(camera *Camera) {
	if r.state == "die" {
		return
	}
	r.drawTarget(camera)
	op := &ebiten.DrawImageOptions{}
	if r.isFlip {
		op.GeoM.Scale(-1, 1)
	}

	op.GeoM.Translate(r.x-r.offsetx, r.y-r.offsety)
	if r.isFlip {
		op.GeoM.Translate(r.size, 0)
	}
	animation := r.animations[r.state]
	camera.DrawImage(animation.GetImage(), op, midgroundLayer)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(r.x-8, r.y-8)
	op.ColorScale.ScaleAlpha(0.5)
	camera.DrawImage(r.shadow, op, midgroundLayer-1)

	if r.isSelected {
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(r.x-8, r.y-8)
		op.ColorScale.ScaleAlpha(0.5)
		camera.DrawImage(r.selection, op, midgroundLayer-50)
	}
}

func (r *Soldier) SetTarget(x, y float64) {
	r.tx = x
	r.ty = y
	r.selectionTimer = SelectionTime
}

func (r *Soldier) SetPosition(x, y float64) {
	r.tx = x
	r.ty = y
	r.x = x
	r.y = y
}

var bloodimages = []string{"blood-1", "blood-2"}

func (r *Soldier) Die(game *Game) {
	if r.state == "die" {
		return
	}
	r.state = "die"

	bloodImage := bloodimages[rand.Intn(2)]
	game.AddDecor(&Decor{
		x:         r.x - r.offsetx,
		y:         r.y - r.offsety,
		z:         forgroundLayer,
		animation: NewAnimation(game.resources.GetImage(bloodImage), 6, 0.1, 20, true),
	})
	game.AddDecor(&Decor{
		x:         r.x - r.offsetx,
		y:         r.y - r.offsety,
		z:         midgroundLayer,
		animation: NewAnimation(game.resources.GetImage(r.unitRes.Die), 4, 0.2, int(r.size), true),
	})
}

func (r *Soldier) drawTarget(camera *Camera) {
	if r.selectionTimer <= 0 {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(r.tx, r.ty)
	op.ColorScale.ScaleAlpha(float32(r.selectionTimer) / SelectionTime)
	camera.DrawImage(r.selection, op, midgroundLayer-50)
}
