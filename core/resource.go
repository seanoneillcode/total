package core

import (
	"os"
	"strings"
	"total/common"

	"github.com/hajimehoshi/ebiten/v2"
)

type Resources struct {
	images        map[string]*ebiten.Image
	unitResources map[string]UnitResource
}

func NewResources() *Resources {
	images := map[string]*ebiten.Image{}
	entries, err := os.ReadDir("res")
	if err != nil {
		panic(err)
	}

	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".png") {
			key, _ := strings.CutSuffix(e.Name(), ".png")
			images[key] = common.LoadImage(e.Name())
		}
	}
	return &Resources{
		images:        images,
		unitResources: unitResources,
	}
}

func (r *Resources) GetImage(id string) *ebiten.Image {
	image, ok := r.images[id]
	if !ok {
		panic("image not found: " + id)
	}
	return image
}

func (r *Resources) GetUnitResource(id string) UnitResource {
	ur, ok := r.unitResources[id]
	if !ok {
		panic("unitresource not found:" + id)
	}
	return ur
}

type UnitResource struct {
	Idle string
	Walk string
	Die  string
	Size int
}

var unitResources = map[string]UnitResource{
	"blue-soldier": {
		Idle: "soldier-idle",
		Walk: "soldier-walk",
		Die:  "soldier-die",
		Size: 16,
	},
	"blue-archer": {
		Idle: "archer-idle-blue",
		Walk: "archer-walk-blue",
		Die:  "archer-die-blue",
		Size: 16,
	},
	"red-soldier": {
		Idle: "soldier-idle",
		Walk: "soldier-walk",
		Die:  "soldier-die",
		Size: 16,
	},
	"red-archer": {
		Idle: "archer-idle",
		Walk: "archer-idle",
		Die:  "archer-idle",
		Size: 16,
	},
	"red-knight": {
		Idle: "horse",
		Walk: "horse",
		Die:  "horse",
		Size: 32,
	},
	"wizard": {
		Idle: "bishop-idle",
		Walk: "bishop-walk",
		Die:  "bishop-die",
		Size: 16,
	},
	"dwarf": {
		Idle: "dwarf-idle",
		Walk: "dwarf-run",
		Die:  "dwarf-idle",
		Size: 24,
	},
	"goblin": {
		Idle: "goblin",
		Walk: "goblin-walk",
		Die:  "goblin-die",
		Size: 16,
	},
	"thug": {
		Idle: "thug",
		Walk: "thug",
		Die:  "thug",
		Size: 16,
	},
}
