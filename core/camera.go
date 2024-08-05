package core

import (
	"sort"
	"total/common"

	"github.com/hajimehoshi/ebiten/v2"
)

const HalfScreenWidth = common.ScreenWidth / 2
const HalfScreenHeight = common.ScreenHeight / 2

const backgroundLayer = 0
const midgroundLayer = 100
const forgroundLayer = 200

type drawCall struct {
	image *ebiten.Image
	op    *ebiten.DrawImageOptions
	y     float64
	z     int
}

type Camera struct {
	x           float64
	y           float64
	screen      *ebiten.Image
	drawCalls   []drawCall
	uiDrawCalls []drawCall
}

func NewCamera() *Camera {
	return &Camera{
		drawCalls:   []drawCall{},
		uiDrawCalls: []drawCall{},
	}
}

func (r *Camera) DrawImage(image *ebiten.Image, op *ebiten.DrawImageOptions, z int) {
	y := op.GeoM.Element(1, 2)
	r.drawCalls = append(r.drawCalls, drawCall{image: image, op: op, y: y, z: z})
}

func (r *Camera) DrawUI(image *ebiten.Image, op *ebiten.DrawImageOptions, z int) {
	r.uiDrawCalls = append(r.uiDrawCalls, drawCall{image: image, op: op, z: z})
}

func (r *Camera) yzSort(i, j int) bool {
	if r.drawCalls[i].z == r.drawCalls[j].z {
		return r.drawCalls[i].y < r.drawCalls[j].y
	}
	return r.drawCalls[i].z < r.drawCalls[j].z
}

func (r *Camera) zSort(i, j int) bool {
	return r.drawCalls[i].z < r.drawCalls[j].z
}

func (r *Camera) Draw() {
	sort.Slice(r.drawCalls, r.yzSort)
	for _, dc := range r.drawCalls {
		dc.op.GeoM.Translate(-r.x+HalfScreenWidth, -r.y+HalfScreenHeight)
		dc.op.GeoM.Scale(common.Scale, common.Scale)
		r.screen.DrawImage(dc.image, dc.op)
	}
	r.drawCalls = nil

	sort.Slice(r.uiDrawCalls, r.zSort)
	for _, dc := range r.uiDrawCalls {
		r.screen.DrawImage(dc.image, dc.op)
	}
	r.uiDrawCalls = nil
}
