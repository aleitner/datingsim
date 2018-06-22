package main

import (
	"fmt"

	"github.com/aleitner/datingsim/datingsim"

	"github.com/hajimehoshi/ebiten"
)

var (
	game *datingsim.Game
)

func update(screen *ebiten.Image) error {
	err := game.Update()
	if err != nil {
		return err
	}

	err = game.Draw(screen)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	title := "Super Date Night Ultra Sunshine Romance 2018!"
	game, err := datingsim.NewGame()
	if err != nil {
		panic(err)
	}
	fmt.Println(game)

	ebiten.Run(update, datingsim.ScreenHeight, datingsim.ScreenWidth, 2, title)
}
