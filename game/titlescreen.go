package game

import (
	"fmt"
	"image/color"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type TitleScene struct {
}

func (s *TitleScene) Update(state *GameState) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		fmt.Println("space button pressed")
		// state.SceneManager.GoTo(NewGameScene())
		return nil
	}
	return nil
}

func (s *TitleScene) Draw(r *ebiten.Image) {
	drawText(r, ScreenWidth/2 - 40, ScreenHeight/4, "Super Date Night Ultra Sunshine Romance 2018!")
	drawText(r, ScreenWidth/2 + 15, ScreenHeight/2, "PRESS SPACE TO START")
}

func drawText(r *ebiten.Image, x, y int, str string) {
	font, _ := truetype.Parse(goregular.TTF)

	text.Draw(r, str, truetype.NewFace(font, nil), x, y, color.White)
}
