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
	decors           []*Decor
}

func NewGame() *Game {
	r := &Game{
		images: map[string]*ebiten.Image{
			// load all the images once, up front
			"player":  common.LoadImage("player.png"),
			"grass-1": common.LoadImage("grass-1.png"),
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

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return common.NormalEscapeError
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	return nil
}

func (r *Game) Draw(screen *ebiten.Image) {
	r.camera.screen = screen
	common.DrawText(screen, "hello", 60, 120)

	for _, d := range r.decors {
		d.Draw(r.camera)
	}

	r.player.Draw(r.camera)

}

func (r *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return common.ScreenWidth * common.Scale, common.ScreenHeight * common.Scale
}

// func (r *Game) drawImage(screen *ebiten.Image, image *ebiten.Image, x float64, y float64) {
// 	op := &ebiten.DrawImageOptions{}
// 	op.GeoM.Translate(x, y)
// 	op.GeoM.Scale(common.Scale, common.Scale)
// 	screen.DrawImage(image, op)
// }
