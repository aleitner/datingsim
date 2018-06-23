package game

import (
	"github.com/hajimehoshi/ebiten"
)

const (
	ScreenWidth  = 420
	ScreenHeight = 600
)

type GameState struct {
	sceneManager *SceneManager
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
