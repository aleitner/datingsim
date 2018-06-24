package game

import (
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten"
)

type GameState struct {
	sceneManager *SceneManager
	Settings     []*Setting
}

// Update updates the current game state.
func (state *GameState) Update() error {
	if state.sceneManager == nil {
		state.sceneManager = &SceneManager{}
		state.sceneManager.GoTo(&TitleScene{})
	}

	if state.Settings == nil {
		ebiten.SetFullscreen(state.Settings[0].Selected != 0)
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
	state.Settings = append(state.Settings, &Setting{Content: "FullScreen", Options: []string{"off", "on"}})
	state.Settings = append(state.Settings, &Setting{Content: "Resolution", Options: []string{"480*360", "1024*768", "1366*768", "1440*900", "1600*900", "1920*1080"}})
}

// Return if the game is fullscreened or not
func (state *GameState) IsFullScreen() bool {
	return state.Settings[0].Selected != 0
}

// Return Width, height of the game
func (state *GameState) Resolution() (width, height int) {
	if state.Settings == nil || len(state.Settings) < 2 {
		return 480, 360
	}

	setting := state.Settings[1]

	slice := strings.Split(setting.Options[setting.Selected], "*")

	width, _ = strconv.Atoi(slice[0])
	height, _ = strconv.Atoi(slice[1])
	return width, height
}
