package game

import (
	"image/color"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func drawText(r *ebiten.Image, x, y int, str string, clr color.Color) {
	font, _ := truetype.Parse(goregular.TTF)

	text.Draw(r, str, truetype.NewFace(font, nil), x, y, clr)
}

// InteractiveElement -- These are all elements that can be interacted with on screen
type InteractiveElement interface {
	Text() string
	Update(*GameState)
}

// BackButton -- Button for returning to previous screen
type BackButton struct {
	text string
	previous Scene
}

func (s *BackButton) Text() string {
	return s.text
}

func (b *BackButton) Update(state *GameState) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		state.sceneManager.GoTo(b.previous)
	}
}

// Setting -- Element for adjusting a setting
type Setting struct {
	text string
	selected int
	options []string
}

func (s *Setting) SelectedOption() string {
	return s.options[s.selected]
}

func (s *Setting) Text() string {
	return s.text
}

func (s *Setting) Update(state *GameState) {

	// Move through options when right key is pressed
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		s.selected += 1
		if s.selected >= len(s.options) {
			s.selected = 0
		}
	}

	// Move through options when left key is pressed
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		s.selected -= 1
		if s.selected < 0 {
			s.selected = len(s.options) - 1
		}
	}
}
