package core

import (
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
	shadow    *ebiten.Image
}

func NewPlayer(game *Game) *Player {
	p := &Player{
		y:         0,
		x:         0,
		speed:     0.5,
		animation: NewAnimation(game.images["player"], 4, 0.2, 16, false),
		shadow:    game.images["unit-shadow"],
	}
	return p
}

func (r *Player) Update(delta float64, game *Game) {
	r.animation.Update(delta, game)
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		r.x = r.x - (r.speed)
		r.isFlip = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		r.x = r.x + (r.speed)
		r.isFlip = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		r.y = r.y - (r.speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		r.y = r.y + (r.speed)
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		mx, my := game.MousePos()
		wx, wy := game.ScreenPosToWorldPos(mx, my)

		if game.selectedUnit == nil {
		exitSelection:
			for _, u := range game.units {
				for _, s := range u.soldiers {
					piSize := s.size / 3
					if common.Overlap(s.x-piSize, s.y-piSize, piSize*2, wx+7, wy+7, 2) {
						game.selectedUnit = u
						game.selectedUnit.GetSelected()
						break exitSelection
					}
				}
			}
		} else {
			game.selectedUnit.MoveTo(wx, wy)
		}
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton1) {
		mx, my := game.MousePos()
		wx, wy := game.ScreenPosToWorldPos(mx, my)
		for _, u := range game.units {
			for _, s := range u.soldiers {
				if common.Overlap(s.x, s.y, 16, wx+4, wy+4, 1) {
					s.Die(game)
				}
			}
		}

	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton2) {
		if game.selectedUnit != nil {
			game.selectedUnit.GetDeSelected()
		}
		game.selectedUnit = nil
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
	camera.DrawImage(r.animation.GetImage(), op, midgroundLayer)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(r.x-1, r.y+6)
	if r.isFlip {
		op.GeoM.Translate(2, 0)
	}
	op.ColorScale.ScaleAlpha(0.5)
	camera.DrawImage(r.shadow, op, midgroundLayer-1)
}
