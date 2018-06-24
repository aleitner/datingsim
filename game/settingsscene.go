package game

import (
	"image/color"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type SettingsScene struct{
	currentElement int
	elements []InteractiveElement
	tempSettings []*Setting
}

func (s *SettingsScene) elementCount() int {
	return len(s.elements) + len(s.tempSettings)
}

func (s *SettingsScene) Update(state *GameState) error {
	if s.tempSettings == nil {
		s.tempSettings = state.settings
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
		s.elements[s.currentElement - len(s.tempSettings)].Update(state)
	} else {
		s.tempSettings[s.currentElement].Update(state)
	}

	return nil
}

func (s *SettingsScene) Draw(state *GameState, screen *ebiten.Image) {
	drawbackground(screen)
	drawText(screen, ScreenWidth/2 + 40, ScreenHeight/10, "Settings", color.Black)

	elementPos := 0

	// Draw Settings
	for _, setting := range s.tempSettings {
		clr := color.Black

		if s.currentElement == elementPos {
			clr = color.White
		}
		drawText(screen, ScreenWidth/3, ScreenHeight/6 + elementPos * 20, setting.Text(), color.Black)
		drawText(screen, ScreenWidth - ScreenWidth/5, ScreenHeight/6 + elementPos * 20, setting.SelectedOption(), clr)

		elementPos += 1
	}

	// Draw other elements element
	for _, element := range s.elements {
		clr := color.Black

		if s.currentElement == elementPos {
			clr = color.White
		}

		drawText(screen, ScreenWidth/3, ScreenHeight/6 + elementPos * 20, element.Text(), clr)

		elementPos += 1
	}
}

// Load All Interactive things into the elements array
// Settings are also clickable but they have an extra method so we'll treat them differently
func (s *SettingsScene) initializeElements(state *GameState) {
	// Load Interactive Save Button
	s.elements = append(s.elements, &SaveSettings{text: "Save", tempSettings: s.tempSettings})
	// Load Interactive Back Button
	s.elements = append(s.elements, &BackButton{text: "Back", previous: &TitleScene{}})
}

// SaveSetting -- Button with pointer to temp settings
type SaveSettings struct {
	text string
	tempSettings []*Setting
}

func (s *SaveSettings) Text() string {
	return s.text
}

// Save the temp Settings
func (s *SaveSettings) Update(state *GameState) {

	// Save the Settings
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		state.settings = s.tempSettings

		// Reload based on the new Settings
		ebiten.SetFullscreen(state.settings[0].selected != 0)
	}
}

func drawbackground(screen *ebiten.Image) {
	i, _ := ebiten.NewImage(ScreenHeight, ScreenWidth, ebiten.FilterDefault)
	i.Fill(&color.RGBA{255, 0, 0, 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(i, op)
}
