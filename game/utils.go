package game

import (
	"bytes"
	"encoding/gob"
	"image/color"
	"log"

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

func copySettings(src []*Setting) (dst []*Setting) {
	var mod bytes.Buffer
	enc := gob.NewEncoder(&mod)
	dec := gob.NewDecoder(&mod)


	err := enc.Encode(src)
	if err != nil {
			log.Fatal("encode error:", err)
	}

	err = dec.Decode(&dst)
	if err != nil {
			log.Fatal("decode error:", err)
	}
	return dst
}

// InteractiveElement -- These are all elements that can be interacted with on screen
type InteractiveElement interface {
	Text() string
	Update(*GameState)
}

// BackButton -- Button for returning to previous screen
type BackButton struct {
	Content string
	previous Scene
}

func (s *BackButton) Text() string {
	return s.Content
}

func (b *BackButton) Update(state *GameState) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		state.sceneManager.GoTo(b.previous)
	}
}

// Setting -- Element for adjusting a setting
type Setting struct {
	Content string
	Selected int
	Options []string
}

func (s *Setting) SelectedOption() string {
	return s.Options[s.Selected]
}

func (s *Setting) Text() string {
	return s.Content
}

func (s *Setting) Update(state *GameState) {

	// Move through options when right key is pressed
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		s.Selected += 1
		if s.Selected >= len(s.Options) {
			s.Selected = 0
		}
	}

	// Move through options when left key is pressed
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		s.Selected -= 1
		if s.Selected < 0 {
			s.Selected = len(s.Options) - 1
		}
	}
}
