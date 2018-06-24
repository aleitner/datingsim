package game

import (
	"github.com/hajimehoshi/ebiten"
)

type Scene interface {
	Update(*GameState) error        // Update the scene by game state
	Draw(*GameState, *ebiten.Image) // Draw onto the Screen that gets passed in
}

type SceneManager struct {
	current  Scene
	previous Scene
}

func (s *SceneManager) Update(state *GameState) error {
	// Do something with scenes
	if s.current == nil {
		return nil
	}

	return s.current.Update(state)
}

func (s *SceneManager) Draw(state *GameState, screen *ebiten.Image) {
	if s.current == nil {
		return
	}

	s.current.Draw(state, screen)
}

func (s *SceneManager) GoTo(scene Scene) {
	if s.current == nil {
		s.previous = s.current
	}

	s.current = scene
}
