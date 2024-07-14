package core

import (
	"fmt"
	"total/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	x         float64
	y         float64
	speed     float64
	animation *Animation
	isFlip    bool
}

func NewPlayer(game *Game) *Player {
	p := &Player{
		y:         0,
		x:         0,
		speed:     30,
		animation: NewAnimation(game.images["player"], 4, 0.2, 16),
	}
	return p
}

func (r *Player) Update(delta float64, game *Game) {
	r.animation.Update(delta, game)
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		r.x = r.x - (delta * r.speed)
		r.isFlip = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		r.x = r.x + (delta * r.speed)
		r.isFlip = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		r.y = r.y - (delta * r.speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		r.y = r.y + (delta * r.speed)
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		fmt.Println("pressed 0")
		mx, my := game.MousePos()
		fmt.Println("pressed 0 ", mx, " ", my)

		if game.selectedSoldier == nil {
			for _, s := range game.soldiers {
				if common.Overlap(s.x, s.y, 16, mx, my, 16) {
					fmt.Println("clicked soldier")
					game.selectedSoldier = s
				}
			}
		} else {
			game.selectedSoldier.tx = mx
			game.selectedSoldier.ty = my
		}
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton1) {
		fmt.Println("pressed 1")
		game.selectedSoldier = nil
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton2) {
		fmt.Println("pressed 2")
		game.selectedSoldier = nil
	}
}

func (r *Player) Draw(camera *Camera) {
	op := &ebiten.DrawImageOptions{}
	if r.isFlip {
		op.GeoM.Scale(-1, 1)
	}
	op.GeoM.Translate(r.x, r.y)
	if r.isFlip {
		op.GeoM.Translate(16, 0)
	}
	camera.DrawImage(r.animation.GetImage(), op)
}
