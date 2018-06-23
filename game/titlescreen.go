package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type TitleScene struct{}

func (s *TitleScene) Update(state *GameState) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.sceneManager.GoTo(&SettingsScene{})
		return nil
	}
	return nil
}

func (s *TitleScene) Draw(r *ebiten.Image) {
	drawText(r, ScreenWidth/2-40, ScreenHeight/4, "Super Date Night Ultra Sunshine Romance 2018!")
	drawText(r, ScreenWidth/2+15, ScreenHeight/2, "PRESS SPACE TO START")
}
