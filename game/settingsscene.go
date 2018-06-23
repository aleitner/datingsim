package game

import (
	"image/color"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

var settings []*Setting

type Setting struct {
	text string
	selected int
	options []string
}

type SettingsScene struct{
	selector int
}

func loadSettings(state *GameState) error {
	settings = append(settings, &Setting{text: "FullScreen", options: []string{"on", "off"}})
	settings = append(settings, &Setting{text: "Resolution", options: []string{"320*480"}})

	return nil
}

func (s *SettingsScene) Update(state *GameState) error {
	if settings == nil {
		loadSettings(state)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		s.selector -= 1
		if s.selector < 0 {
			s.selector = len(settings) - 1
		}
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		s.selector += 1
		if s.selector >= len(settings) {
			s.selector = 0
		}
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		settings[s.selector].selected += 1
		if settings[s.selector].selected >= len(settings[s.selector].options) {
			settings[s.selector].selected = 0
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.sceneManager.GoTo(&TitleScene{})
		return nil
	}

	return nil
}

func (s *SettingsScene) Draw(r *ebiten.Image) {
	drawbackground(r)
	drawText(r, ScreenWidth/2 + 40, ScreenHeight/10, "Settings", color.Black)

	if settings == nil {
		return
	}

	for i, v := range settings {
		clr := color.Black

		if s.selector == i {
			clr = color.White
		}
		drawText(r, ScreenWidth/3, ScreenHeight/6 + i * 20, v.text, color.Black)
		drawText(r, ScreenWidth - ScreenWidth/5, ScreenHeight/6 + i * 20, v.options[v.selected], clr)
	}

}

func drawbackground(screen *ebiten.Image) {
	i, _ := ebiten.NewImage(ScreenHeight, ScreenWidth, ebiten.FilterDefault)
	i.Fill(&color.RGBA{255, 0, 0, 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(i, op)
}
