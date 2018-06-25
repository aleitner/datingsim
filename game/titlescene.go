package game

import (
	"log"
	"os"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"image/png"
	"image/color"
)

type TitleScene struct{
	background *Background
}

func (s *TitleScene) Update(state *GameState) error {
	if s.background == nil {
		var err error
		s.background, err = NewBackground("/Users/alexleitner/go/src/github.com/aleitner/datingsim/game/assets/title.png")
		if err != nil {
			log.Fatal(err)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.sceneManager.GoTo(&GameSettingsScene{})
		return nil
	}
	return nil
}

func (s *TitleScene) Draw(state *GameState, screen *ebiten.Image) {
	width, height := state.Resolution()

	if s.background != nil {
		s.background.Draw(screen)
	}

	drawText(screen, width/2+15, height/2, "PRESS SPACE TO START", color.White)
}

type Background struct {
	img *ebiten.Image
}

func NewBackground(path string) (*Background, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	img2, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}

	return &Background{img2}, nil
}

func (b *Background) Draw(screen *ebiten.Image) {
	if b.img == nil {
		return
	}
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(b.img, op)
}
