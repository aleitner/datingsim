package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type GameSettingsScene struct {
	currentElement int
	elements       []InteractiveElement
	tempSettings   []*Setting
}

func (s *GameSettingsScene) elementCount() int {
	return len(s.elements) + len(s.tempSettings)
}

func (s *GameSettingsScene) Update(state *GameState) error {
	if s.tempSettings == nil {
		s.tempSettings = copySettings(state.Settings)
	}

	if len(s.elements) <= 0 {
		s.initializeElements(state)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		s.currentElement -= 1
		if s.currentElement < 0 {
			s.currentElement = s.elementCount() - 1
		}
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		s.currentElement += 1
		if s.currentElement >= s.elementCount() {
			s.currentElement = 0
		}
		return nil
	}

	// Perform Update for the current Element that is Highlighted
	// Settings and Elements are separated so we have to determine
	// if the settings are highlighted or if the other elements are highlighted
	if s.currentElement >= len(s.tempSettings) {
		s.elements[s.currentElement-len(s.tempSettings)].Update(state)
	} else {
		s.tempSettings[s.currentElement].Update(state)
	}

	return nil
}

func (s *GameSettingsScene) Draw(state *GameState, screen *ebiten.Image) {
	width, height := state.Resolution()

	drawbackground(screen, width, height)
	drawText(screen, width/10, height/10, "Settings", color.Black)

	elementPos := 0

	// Draw Settings
	for _, setting := range s.tempSettings {
		setting.Draw(screen, width*2/5, height/6+elementPos*20, s.currentElement == elementPos)

		elementPos += 1
	}

	// Draw other elements element
	for _, element := range s.elements {
		element.Draw(screen, width*2/5, height/6+elementPos*20, s.currentElement == elementPos)

		elementPos += 1
	}
}

// Load All Interactive things into the elements array
// Settings are also clickable but they have an extra method so we'll treat them differently
func (s *GameSettingsScene) initializeElements(state *GameState) {
	// Load Interactive Save Button
	s.elements = append(s.elements, &SaveSettings{Content: "Save", tempSettings: s.tempSettings})
	// Load Interactive Back Button
	s.elements = append(s.elements, &BackButton{Content: "Back", previous: &TitleScene{}})
}

// SaveSetting -- Button with pointer to temp settings
type SaveSettings struct {
	Content      string
	tempSettings []*Setting
}

func (s *SaveSettings) Text() string {
	return s.Content
}

// Save the temp Settings
func (s *SaveSettings) Update(state *GameState) {

	// Save the Settings
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		state.Settings = copySettings(s.tempSettings)

		if len(state.Settings) <= 0 {
			return
		}

		// Reload based on the new Settings
		ebiten.SetFullscreen(state.Settings[0].Selected != 0)
		ebiten.SetScreenSize(state.Resolution())

		// TODO: Write settings to fs
	}
}

func (s *SaveSettings) Draw(screen *ebiten.Image, x, y int, isSelected bool) {
	clr := color.Black

	if isSelected {
		clr = color.White
	}

	drawText(screen, x, y, s.Content, clr)
}
