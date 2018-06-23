package game

import (
	"github.com/hajimehoshi/ebiten"
)

var (
	ScreenWidth  = 360
	ScreenHeight = 480
)

type Setting struct {
	text string
	selected int
	options []string
}

type GameState struct {
	sceneManager *SceneManager
	settings []*Setting
}

// Update updates the current game state.
func (state *GameState) Update() error {
	if state.sceneManager == nil {
		state.sceneManager = &SceneManager{}
		state.sceneManager.GoTo(&TitleScene{})
	}

	if state.settings == nil {
		state.LoadDefaultSettings()
	}

	state.sceneManager.current.Update(state)
	return nil
}

func (state *GameState) Draw(screen *ebiten.Image) error {
	if state.sceneManager == nil {
		return nil
	}

	state.sceneManager.Draw(state, screen)
	return nil
}

func (state *GameState) LoadDefaultSettings() {
	state.settings = append(state.settings, &Setting{text: "FullScreen", options: []string{"on", "off"}})
	state.settings = append(state.settings, &Setting{text: "Resolution", options: []string{"320*480"}})
}
