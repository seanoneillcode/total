package core

import (
	"time"
	"total/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	lastUpdateCalled time.Time
	player           *Player
	images           map[string]*ebiten.Image
	camera           *Camera
	cursor           *Cursor
	decors           []*Decor
	soldiers         []*Soldier
	selectedSoldier  *Soldier
}

func NewGame() *Game {
	r := &Game{
		images: map[string]*ebiten.Image{
			// load all the images once, up front
			"player":       common.LoadImage("player.png"),
			"grass-1":      common.LoadImage("grass-1.png"),
			"cursor":       common.LoadImage("cursor.png"),
			"soldier-idle": common.LoadImage("soldier-idle.png"),
			"soldier-walk": common.LoadImage("soldier-walk.png"),
		},
		lastUpdateCalled: time.Now(),
		camera:           &Camera{},
	}
	r.player = NewPlayer(r)
	r.decors = []*Decor{
		{
			x:         0,
			y:         0,
			animation: NewAnimation(r.images["grass-1"], 4, 0.2, 16),
		},
		{
			x:         64,
			y:         0,
			animation: NewAnimation(r.images["grass-1"], 4, 0.2, 16),
		},
		{
			x:         0,
			y:         32,
			animation: NewAnimation(r.images["grass-1"], 4, 0.2, 16),
		},
	}
	r.cursor = NewCursor(r)
	r.soldiers = []*Soldier{
		NewSoldier(r, -120, 20),
		NewSoldier(r, 20, 120),
		NewSoldier(r, 120, 120),
		NewSoldier(r, 20, 0),
	}

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
	for _, s := range r.soldiers {
		s.Update(delta, r)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return common.NormalEscapeError
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	r.cursor.Update()

	return nil
}

func (r *Game) Draw(screen *ebiten.Image) {
	r.camera.screen = screen
	common.DrawText(screen, "hello", 60, 120)

	for _, d := range r.decors {
		d.Draw(r.camera)
	}
	for _, s := range r.soldiers {
		s.Draw(r.camera)
	}

	r.player.Draw(r.camera)
	r.cursor.Draw(r.camera)
}

func (r *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return common.ScreenWidth * common.Scale, common.ScreenHeight * common.Scale
}

func (r *Game) MousePos() (float64, float64) {
	x, y := ebiten.CursorPosition()
	return float64(x/common.Scale) - 8, float64(y/common.Scale) - 8
}
