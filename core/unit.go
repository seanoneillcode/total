package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"math"
	"total/common"
)

const SelectionTime = 1.0

type Unit struct {
	x            float64
	y            float64
	soldiers     []*Soldier
	unitTemplate *UnitTemplate
	bannerTop    *Decor
	bannerBottom *Decor
}

func NewUnit(game *Game) *Unit {
	top := ebiten.NewImageFromImage(game.resources.GetImage("banner").SubImage(image.Rect(0, 0, 64, 16)))
	bottom := ebiten.NewImageFromImage(game.resources.GetImage("banner").SubImage(image.Rect(0, 16, 64, 32)))
	topAnim := NewAnimation(top, 4, 0.4, 16, false)
	bottomAnim := NewAnimation(bottom, 4, 0.4, 16, false)
	bottomAnim.frame = topAnim.frame
	return &Unit{
		soldiers: []*Soldier{},
		bannerTop: &Decor{
			z:         midgroundLayer + 10,
			animation: topAnim,
		},
		bannerBottom: &Decor{
			z:         midgroundLayer,
			animation: bottomAnim,
		},
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
	u.bannerTop.x = x
	u.bannerTop.y = y - 32 + 8
	u.bannerBottom.x = x
	u.bannerBottom.y = y - 8
}

func (u *Unit) Update(delta float64, game *Game) {
	needsClearing := false
	for _, s := range u.soldiers {
		s.Update(delta, game)
		if s.state == "die" {
			needsClearing = true
		}
	}

	if needsClearing {
		newSlice := []*Soldier{}
		for _, s := range u.soldiers {
			if s.state != "die" {
				newSlice = append(newSlice, s)
			}
		}
		u.soldiers = newSlice
	}
	u.bannerTop.animation.Update(delta, game)
	u.bannerBottom.animation.Update(delta, game)
}

func (u *Unit) Draw(camera *Camera) {
	for _, s := range u.soldiers {
		s.Draw(camera)
	}
	u.bannerTop.Draw(camera)
	u.bannerBottom.Draw(camera)
}

type UnitTemplate struct {
	pos map[int]Soldier
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
