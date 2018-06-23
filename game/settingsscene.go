package game

import (
	"image/color"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type SettingsScene struct{
	selector int
}

func (s *SettingsScene) Update(state *GameState) error {
	settings := state.settings

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
		setting := settings[s.selector]
		setting.selected += 1
		if setting.selected >= len(setting.options) {
			setting.selected = 0
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.sceneManager.GoTo(&TitleScene{})
		return nil
	}

	return nil
}

func (s *SettingsScene) Draw(state *GameState, screen *ebiten.Image) {
	drawbackground(screen)
	drawText(screen, ScreenWidth/2 + 40, ScreenHeight/10, "Settings", color.Black)

	for i, v := range state.settings {
		clr := color.Black

		if s.selector == i {
			clr = color.White
		}
		drawText(screen, ScreenWidth/3, ScreenHeight/6 + i * 20, v.text, color.Black)
		drawText(screen, ScreenWidth - ScreenWidth/5, ScreenHeight/6 + i * 20, v.options[v.selected], clr)
	}

}

func drawbackground(screen *ebiten.Image) {
	i, _ := ebiten.NewImage(ScreenHeight, ScreenWidth, ebiten.FilterDefault)
	i.Fill(&color.RGBA{255, 0, 0, 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(i, op)
}
