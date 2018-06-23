package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type SettingsScene struct{}

func (s *SettingsScene) Update(state *GameState) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.sceneManager.GoTo(&TitleScene{})
		return nil
	}

	return nil
}

func (s *SettingsScene) Draw(r *ebiten.Image) {
	drawText(r, ScreenWidth/2-40, ScreenHeight/4, "Settings")
}
