package main

import (
	"errors"
	"log"
	"total/common"
	"total/core"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := core.NewGame()

	ebiten.SetWindowSize(common.ScreenWidth*common.Scale, common.ScreenHeight*common.Scale)
	ebiten.SetWindowTitle("Total")
	err := ebiten.RunGame(game)
	if err != nil {
		if errors.Is(err, common.NormalEscapeError) {
			log.Println("exiting normally")
		} else {
			log.Fatal(err)
		}
	}
}
