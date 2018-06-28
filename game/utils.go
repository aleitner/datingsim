package game

import (
	"image/color"
	"strconv"
	"strings"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

type Scale func() (int, int)

func drawText(r *ebiten.Image, x, y int, str string, clr color.Color) {
	font, _ := truetype.Parse(goregular.TTF)

	text.Draw(r, str, truetype.NewFace(font, nil), x, y, clr)
}

func drawbackground(screen *ebiten.Image, width, height int) {
	i, _ := ebiten.NewImage(width, height, ebiten.FilterDefault)
	i.Fill(&color.RGBA{255, 0, 0, 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(i, op)
}

// IsFullScreen -- Return if the game is fullscreened or not
func IsFullScreen() bool {
	return GameSettings[FullScreenSetting].SelectedOption() == "on"
}

// Resolution -- Return Width, height of the game
func Resolution() (width, height int) {
	slice := strings.Split(GameSettings[ResolutionSetting].SelectedOption(), "*")

	width, _ = strconv.Atoi(slice[0])
	height, _ = strconv.Atoi(slice[1])
	return width, height
}
