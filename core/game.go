package core

import (
	"image/color"
	"math/rand"
	"time"
	"total/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	lastUpdateCalled time.Time
	player           *Player
	camera           *Camera
	cursor           *Cursor
	decors           []*Decor
	units            []*Unit
	selectedUnit     *Unit
	resources        *Resources
}

func NewGame() *Game {
	r := &Game{
		lastUpdateCalled: time.Now(),
		camera:           NewCamera(),
		resources:        NewResources(),
	}
	r.player = NewPlayer(r)
	r.decors = []*Decor{}
	grassImageNames := []string{"grass-1", "grass-2", "grass-3", "grass-4"}
	for i := 0; i < 64; i++ {
		r.decors = append(r.decors, &Decor{
			x:         float64(rand.Intn(common.ScreenWidth*2) - common.ScreenWidth),
			y:         float64(rand.Intn(common.ScreenHeight*2) - common.ScreenHeight),
			animation: NewAnimation(r.resources.GetImage(grassImageNames[rand.Intn(4)]), 4, 0.2, 16, false),
		})
	}
	r.cursor = NewCursor(r)
	r.units = []*Unit{
		NewUnit(r),
		NewUnit(r),
		NewUnit(r),
		NewUnit(r),
	}
	for i := 0; i < 8; i++ {
		r.units[0].AddSoldier(NewSoldier(r, 0, 0, "blue-soldier"))
	}
	for i := 0; i < 18; i++ {
		r.units[1].AddSoldier(NewSoldier(r, 0, 0, "blue-archer"))
	}
	for i := 0; i < 4; i++ {
		r.units[2].AddSoldier(NewSoldier(r, 0, 0, "red-knight"))
	}
	for i := 0; i < 2; i++ {
		r.units[3].AddSoldier(NewSoldier(r, 0, 0, "wizard"))
	}

	r.units[0].MoveTo(-40, -60)
	r.units[1].MoveTo(-60, 40)
	r.units[2].MoveTo(30, 30)
	r.units[3].MoveTo(50, -10)

	return r
}

func (r *Game) Update() error {
	delta := float64(time.Now().Sub(r.lastUpdateCalled).Milliseconds()) / 1000
	r.lastUpdateCalled = time.Now()

	r.player.Update(delta, r)
	// set the camera to follow the player
	r.camera.x = r.player.x
	r.camera.y = r.player.y

	for _, d := range r.decors {
		d.animation.Update(delta, r)
	}
	for _, u := range r.units {
		u.Update(delta, r)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return common.NormalEscapeError
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	r.cursor.Update(delta, r)

	return nil
}

func (r *Game) Draw(screen *ebiten.Image) {
	// 32, 74, 26
	backgroundColor := color.RGBA{
		R: 32,
		G: 74,
		B: 26,
		A: 255,
	}
	screen.Fill(backgroundColor)
	r.camera.screen = screen
	common.DrawText(screen, "tactics game prototype", 16, 8)

	for _, d := range r.decors {
		d.Draw(r.camera)
	}
	for _, u := range r.units {
		u.Draw(r.camera)
	}

	r.player.Draw(r.camera)
	r.cursor.Draw(r.camera)

	r.camera.Draw()
}

func (r *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return common.ScreenWidth * common.Scale, common.ScreenHeight * common.Scale
}

func (r *Game) MousePos() (float64, float64) {
	x, y := ebiten.CursorPosition()
	return float64(x/common.Scale) - 8, float64(y/common.Scale) - 8
}

func (r *Game) ScreenPosToWorldPos(x float64, y float64) (float64, float64) {
	return x + r.camera.x - HalfScreenWidth, y + r.camera.y - HalfScreenHeight
}

func (r *Game) AddDecor(decor *Decor) {
	r.decors = append(r.decors, decor)
}
