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
	stats            *Stats
}

func NewGame() *Game {
	r := &Game{
		lastUpdateCalled: time.Now(),
		camera:           NewCamera(),
		resources:        NewResources(),
		stats:            NewStats(),
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
		NewUnit(r, "blue-soldier", 8, 12),
		NewUnit(r, "blue-archer", 18, 20),
		NewUnit(r, "red-knight", 4, 24),
		NewUnit(r, "wizard", 2, 16),
		NewUnit(r, "dwarf", 16, 16),
		NewUnit(r, "goblin", 48, 10),
	}

	r.units[0].SetPosition(-40, -60)
	r.units[1].SetPosition(-60, 40)
	r.units[2].SetPosition(30, 30)
	r.units[3].SetPosition(50, -10)
	r.units[4].SetPosition(200, 0)
	r.units[5].SetPosition(200, 60)

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
		R: 25,
		G: 60,
		B: 62,
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
