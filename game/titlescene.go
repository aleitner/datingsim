package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"image/color"
)

type TitleScene struct{}

func (s *TitleScene) Update(state *GameState) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.sceneManager.GoTo(&GameSettingsScene{})
		return nil
	}
	return nil
}

func (s *TitleScene) Draw(state *GameState, r *ebiten.Image) {
	width, height := state.Resolution()
	drawText(r, width/2-40, height/4, "Super Date Night Ultra Sunshine Romance 2018!", color.White)
	drawText(r, width/2+15, height/2, "PRESS SPACE TO START", color.White)
}
