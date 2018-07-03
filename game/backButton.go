package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// BackButton -- Button for returning to previous screen
type BackButton struct {
	content     string
	previous    Scene
	highlighted bool
	scalefn     Scale
}

func NewBackButton(previous Scene, scale Scale) *BackButton {
	return &BackButton{content: "Back", previous: previous, scalefn: scale}
}

func (b *BackButton) Highlight(highlighted bool) {
	b.highlighted = highlighted
}

func (b *BackButton) Update(state *GameState) {
	if b.highlighted == false {
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		state.sceneManager.GoTo(b.previous)
	}
}

func (b *BackButton) Draw(screen *ebiten.Image) {
	x, y := b.scalefn()

	clr := color.Black

	if b.highlighted {
		clr = color.White
	}

	drawText(screen, x, y, b.content, clr)
}
