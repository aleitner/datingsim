package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type GameSettingsScene struct {
	initialized     bool
	currentElement  int
	previousElement int
	tempSettings    []*Setting
	backButton      *BackButton
	saveButton      *SaveSettings
}

// Initialize the GameSettingsScene values
func (s *GameSettingsScene) initialize(state *GameState) {
	s.backButton = NewBackButton(
		state.sceneManager.previous,
		func() (int, int) { width, height := Resolution(); return width * 9 / 10, height * 9 / 10 })

	s.currentElement = 0
	s.previousElement = 1 // Set this to 1 so that there is no conflict with currentElement
	s.tempSettings = createTempSettings(GameSettings)
	s.saveButton = NewSaveSettingsButton(func() (int, int) { width, height := Resolution(); return width * 1 / 10, height * 9 / 10 })

	s.initialized = true
}

func (s *GameSettingsScene) elementCount() int {
	// Back Button, Save Button, Settings
	return 2 + len(s.tempSettings)
}

func (s *GameSettingsScene) Update(state *GameState) error {
	if s.initialized == false {
		s.initialize(state)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		s.previousElement = s.currentElement
		s.currentElement -= 1
		if s.currentElement < 0 {
			s.currentElement = s.elementCount() - 1
		}
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		s.previousElement = s.currentElement
		s.currentElement += 1
		if s.currentElement >= s.elementCount() {
			s.currentElement = 0
		}
		return nil
	}

	for _, setting := range s.tempSettings {
		setting.Update(state)
	}

	s.backButton.Update(state)
	s.saveButton.Update(state, s.tempSettings)

	// Highlight new and unhighlight old
	s.highlightOption()

	return nil
}

func (s *GameSettingsScene) Draw(screen *ebiten.Image) {
	gw, gh := Resolution()

	drawbackground(screen, gw, gh)
	drawText(screen, gw/10, gh/10, "Settings", color.Black)

	// Draw Settings
	for _, setting := range s.tempSettings {
		setting.Draw(screen)
	}

	s.backButton.Draw(screen)
	s.saveButton.Draw(screen)
}

// Highlight the currently selected Option
func (s *GameSettingsScene) highlightOption() {
	// Highlight of settings
	for i, setting := range s.tempSettings {
		if s.previousElement == i {
			setting.Highlight(false)
		} else if s.currentElement == i {
			setting.Highlight(true)
		}
	}

	// Highlight of backbutton
	if s.currentElement == len(s.tempSettings) {
		s.backButton.Highlight(true)
	} else if s.previousElement == len(s.tempSettings) {
		s.backButton.Highlight(false)
	}

	// Highlight of saveButton
	if s.currentElement == len(s.tempSettings)+1 {
		s.saveButton.Highlight(true)
	} else if s.previousElement == len(s.tempSettings)+1 {
		s.saveButton.Highlight(false)
	}
}

// Create the temp settings
func createTempSettings(settings []*Setting) (tempSettings []*Setting) {

	for i, s := range settings {
		scalefn := initScaleClosure(i)
		tempSettings = append(
			tempSettings,
			&Setting{
				content:        s.content,
				selectedOption: s.selectedOption,
				options:        s.options,
				scalefn:        scalefn})
	}

	return tempSettings
}

// Settings scale functions need yoffset
func initScaleClosure(i int) func() (int, int) {
	yoffset := i
	// closure. function has access to text even after exiting this block
	return func() (int, int) {
		gw, gh := Resolution()
		return gw * 2 / 5, gh/6 + yoffset*20
	}
}
