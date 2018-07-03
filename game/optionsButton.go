package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// OptionsButton -- Button for returning to previous screen
type OptionsButton struct {
	content     string
	highlighted bool
	scalefn     Scale
}

func NewOptionsButton(scale Scale) *OptionsButton {
	return &OptionsButton{content: "Settings", scalefn: scale}
}

func (b *OptionsButton) Highlight(highlighted bool) {
	b.highlighted = highlighted
}

func (b *OptionsButton) Update(state *GameState) {
	if b.highlighted == false {
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		state.sceneManager.GoTo(&GameSettingsScene{})
	}
}

func (b *OptionsButton) Draw(screen *ebiten.Image) {
	x, y := b.scalefn()

	clr := color.Black

	if b.highlighted {
		clr = color.White
	}

	drawText(screen, x, y, b.content, clr)
}
