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

	fmt.Printf("%+v\n", gamestate)

	ebiten.Run(update, game.ScreenHeight, game.ScreenWidth, 1, title)
}
