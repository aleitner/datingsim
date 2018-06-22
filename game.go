package datingsim

import (
        "github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	ScreenWidth  = 420
	ScreenHeight = 600
)

type Game struct {
}

func NewGame() (*Game, error) {
	g := &Game{}
	return g, nil
}

// Update updates the current game state.
func (g *Game) Update() error {
	return nil
}

func (g* Game) Draw(screen *ebiten.Image) error {
  ebitenutil.DebugPrint(screen, "Hello world!")
  return nil
}
