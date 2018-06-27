package game

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/nfnt/resize"
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
		s.background.Draw(screen, width, height)
	}

	drawText(screen, width/2+15, height/2, "PRESS SPACE TO START", color.White)
}

type Background struct {
	img image.Image
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

	return &Background{img}, nil
}

func (b *Background) Draw(screen *ebiten.Image, width, height int) {
	if b.img == nil {
		return
	}

	// Resize the image to match the resolution
	resizedIMG := resize.Resize(uint(width), 0, b.img, resize.Lanczos3)

	// Convert the imageto ebiten format
	img, err := ebiten.NewImageFromImage(resizedIMG, ebiten.FilterDefault)
	if err != nil {
		return
	}

	_, ih := img.Size()
	op := &ebiten.DrawImageOptions{}

	// translate the image to the proper y coordinate if the image is shorter than the height
	if heightDiff := height - ih; heightDiff > 0 {
		op.GeoM.Translate(0, float64(heightDiff/2))
	}

	screen.DrawImage(img, op)
}
