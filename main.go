package main

import (
	"fmt"

	"github.com/aleitner/datingsim/game"

	"github.com/hajimehoshi/ebiten"
)

var gamestate *game.GameState

func update(screen *ebiten.Image) error {
	err := gamestate.Update()
	if err != nil {
		return err
	}

	err = gamestate.Draw(screen)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Starting...")

	title := "Super Date Night Ultra Sunshine Romance 2018!"
	gamestate = &game.GameState{}
	gamestate.LoadDefaultSettings()
	ebiten.SetFullscreen(gamestate.IsFullScreen())
	fmt.Println(gamestate.IsFullScreen())
	width, height := gamestate.Resolution()
	fmt.Println(width, height)
	// Load Settings into gamestate.Settings if they already exist

	fmt.Printf("%+v\n", gamestate)

	ebiten.Run(update, width, height, 1, title)
}
