package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// SaveSetting -- Button with pointer to temp settings
type SaveSettings struct {
	content     string
	highlighted bool
	scalefn     Scale
}

func NewSaveSettingsButton(scale Scale) *SaveSettings {
	return &SaveSettings{content: "Save", scalefn: scale}
}

func (s *SaveSettings) Highlight(highlighted bool) {
	s.highlighted = highlighted
}

// Save the temp Settings
func (s *SaveSettings) Update(state *GameState, tempSettings []*Setting) {

	if s.highlighted != true {
		return
	}

	// Save the Settings
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		for i, tmp := range tempSettings {
			GameSettings[i].selectedOption = tmp.selectedOption
		}

		// Reload based on the new Settings
		ebiten.SetFullscreen(IsFullScreen())
		ebiten.SetScreenSize(Resolution())

		// TODO: Write settings to fs
	}
}

func (s *SaveSettings) Draw(screen *ebiten.Image) {
	clr := color.Black

	if s.highlighted {
		clr = color.White
	}

	x, y := s.scalefn()
	drawText(screen, x, y, s.content, clr)
}
