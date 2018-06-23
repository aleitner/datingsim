package game

import (
	"github.com/hajimehoshi/ebiten"
)

var (
	ScreenWidth  = 360
	ScreenHeight = 480
)

type GameState struct {
	sceneManager *SceneManager
	settings []Setting
}

// Update updates the current game state.
func (state *GameState) Update() error {
	if state.sceneManager == nil {
		state.sceneManager = &SceneManager{}
		state.sceneManager.GoTo(&TitleScene{})
	}

	state.sceneManager.current.Update(state)
	return nil
}

func (state *GameState) Draw(screen *ebiten.Image) error {
	if state.sceneManager == nil {
		return nil
	}

	state.sceneManager.Draw(screen)
	return nil
}
