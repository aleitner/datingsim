package game

import (
	// "fmt"
	"reflect"
	"image/color"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type SettingsScene struct{
	currentElement int
	elements []InteractiveElement
	tempSettings []*Setting
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
			s.currentElement = len(s.elements) - 1
		}
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		s.currentElement += 1
		if s.currentElement >= len(s.elements) {
			s.currentElement = 0
		}
		return nil
	}

	s.elements[s.currentElement].Action(state)

	return nil
}

func (s *SettingsScene) Draw(state *GameState, screen *ebiten.Image) {
	drawbackground(screen)
	drawText(screen, ScreenWidth/2 + 40, ScreenHeight/10, "Settings", color.Black)

	// Load Each element
	for i, v := range s.elements {
		clr := color.Black

		if s.currentElement == i {
			clr = color.White
		}
		drawText(screen, ScreenWidth/3, ScreenHeight/6 + i * 20, v.Text(), clr)
		if reflect.TypeOf(v) == reflect.TypeOf(&Setting{}) {
			drawText(screen, ScreenWidth - ScreenWidth/5, ScreenHeight/6 + i * 20, s.tempSettings[i].SelectedOption(), color.Black)
		}
	}
}

// Load All Interactive things into the elements array
func (s *SettingsScene) initializeElements(state *GameState) {

	// Load all the Interactive Settings
	for _, setting := range s.tempSettings {
		s.elements = append(s.elements, setting)
	}

	// Load Interactive Back Button
	s.elements = append(s.elements, &BackButton{text: "Back", previous: &TitleScene{}})

	// Load Interactive Save Button
	s.elements = append(s.elements, &SaveSettings{text: "Save", tempSettings: s.tempSettings})
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
func (s *SaveSettings) Action(state *GameState) {

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
