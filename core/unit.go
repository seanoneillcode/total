package core

import (
	"math"
	"total/common"

	"github.com/hajimehoshi/ebiten/v2"
)

const SelectionTime = 1.0

type Unit struct {
	x              float64
	y              float64
	soldiers       []*Soldier
	unitTemplate   *UnitTemplate
	targetMarker   *Animation
	selectionTimer float64
}

func NewUnit(game *Game) *Unit {
	return &Unit{
		targetMarker: NewAnimation(game.images["selection"], 2, 0.2, 16),
		soldiers:     []*Soldier{},
	}
}

func (u *Unit) AddSoldier(soldier *Soldier) {
	u.soldiers = append(u.soldiers, soldier)
}

// set the template as the target
func (u *Unit) MoveTo(x, y float64) {
	if x == u.x && y == u.y {
		return
	}
	// calculate the direction
	dirx := x - u.x
	diry := y - u.y

	// normalize it
	dirx, diry = common.Normalize(dirx, diry)

	// get right angle
	rxdir := -diry
	rydir := dirx

	total := len(u.soldiers)
	// part := total / 3.0
	row := int(math.Ceil(math.Sqrt(float64(total))))
	column := row

	// target pos is the top left, should be middle
	px := x - ((float64(column/2) - 0.5) * rxdir * 16)
	py := y - ((float64(column/2) - 0.5) * rydir * 16)
	tx := px
	ty := py

	index := 0
	// the unit is square for now
	for i := 0; i < row && index < total; i++ {
		tx = px + (float64(i) * (16 * -dirx))
		ty = py + (float64(i) * (16 * -diry))
		for j := 0; j < column && index < total; j++ {
			u.soldiers[index].SetTarget(tx, ty)
			tx = tx + (rxdir * 16)
			ty = ty + (rydir * 16)
			index++
		}
	}
	u.x = x
	u.y = y
	u.selectionTimer = SelectionTime
}

func (u *Unit) Update(delta float64, game *Game) {
	for _, s := range u.soldiers {
		s.Update(delta, game)
	}
	if u.selectionTimer > 0 {
		u.selectionTimer = u.selectionTimer - delta
	}
}

func (u *Unit) Draw(camera *Camera) {
	for _, s := range u.soldiers {
		s.Draw(camera)
		u.drawSelection(camera, s.tx, s.ty)
	}
}

type UnitTemplate struct {
	pos map[int]Soldier
}

func (u *Unit) drawSelection(camera *Camera, x float64, y float64) {
	if u.selectionTimer <= 0 {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleAlpha(float32(u.selectionTimer) / SelectionTime)
	camera.DrawImage(u.targetMarker.GetImage(), op, midgroundLayer-50)
}

func (u *Unit) GetSelected() {
	for _, s := range u.soldiers {
		s.isSelected = true
	}
}

func (u *Unit) GetDeSelected() {
	for _, s := range u.soldiers {
		s.isSelected = false
	}
}
