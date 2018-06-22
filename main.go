package main

import (
        "github.com/aleitner/datingsim/game"

        "github.com/hajimehoshi/ebiten"
        "github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	game *datingsim.Game
)

func update(screen *ebiten.Image) error {
        err := game.Update()
        game.Draw()
        return nil
}

func main() {
        title := "Super Date Night Ultra Sunshine Romance 2018!"
        game, err := datingsim.NewGame()
        if err != nil {
          return err
        }

        ebiten.Run(update, game.ScreenWidth, game.ScreenHeight, 2, title)
}
