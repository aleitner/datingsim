package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// Setting -- Element for adjusting a setting
type Setting struct {
	content        string
	selectedOption int
	options        []string
	highlighted    bool
	scalefn        Scale
}

func (s *Setting) SelectedOption() string {
	return s.options[s.selectedOption]
}

func (s *Setting) Highlight(highlighted bool) {
	s.highlighted = highlighted
}

func (s *Setting) Update(state *GameState) {
	if s.highlighted != true {
		return
	}

	// Move through options when right key is pressed
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		s.selectedOption += 1
		if s.selectedOption >= len(s.options) {
			s.selectedOption = 0
		}
	}

	// Move through options when left key is pressed
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		s.selectedOption -= 1
		if s.selectedOption < 0 {
			s.selectedOption = len(s.options) - 1
		}
	}
}

func (s *Setting) Draw(screen *ebiten.Image) {
	x, y := s.scalefn()

	clr := color.Black

	if s.highlighted {
		clr = color.White
	}

	drawText(screen, x, y, s.content, color.Black)
	drawText(screen, x+100, y, s.SelectedOption(), clr)
}
