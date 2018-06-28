package game

import (
	"log"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type TitleScene struct {
	background *Background
}

func (s *TitleScene) Update(state *GameState) error {

	if s.background == nil {
		var err error
		s.background, err = NewBackground("/Users/alexleitner/go/src/github.com/aleitner/datingsim/game/assets/title.png")
		if err != nil {
			log.Fatal(err)
		}
		
	} else {
		s.background.Update(state)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.sceneManager.GoTo(&GameSettingsScene{})
		return nil
	}
	return nil
}

func (s *TitleScene) Draw(screen *ebiten.Image) {
	if s.background != nil {
		s.background.Draw(screen)
	}

	gw, gh := Resolution()

	drawText(screen, gw/2, gh/2, "PRESS SPACE TO START", color.Black)
}
