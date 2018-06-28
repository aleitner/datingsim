package game

import (
	"github.com/hajimehoshi/ebiten"
)

type GameState struct {
	sceneManager *SceneManager
}

// Update -- updates the current game state.
func (state *GameState) Update() error {

	// Initialize SceneManager and go to Title Screen
	// This happens when starting a new game
	if state.sceneManager == nil {
		state.sceneManager = NewSceneManager()
		state.sceneManager.GoTo(&TitleScene{})
	}

	// Update the current Scene
	state.sceneManager.Update(state)
	return nil
}

// Draw -- Draw the current Scene
func (state *GameState) Draw(screen *ebiten.Image) error {
	if state.sceneManager == nil {
		return nil
	}

	state.sceneManager.Draw(state, screen)
	return nil
}
