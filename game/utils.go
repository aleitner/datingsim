package game

import (
	"image/color"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

func drawText(r *ebiten.Image, x, y int, str string, clr color.Color) {
	font, _ := truetype.Parse(goregular.TTF)

	text.Draw(r, str, truetype.NewFace(font, nil), x, y, clr)
}
