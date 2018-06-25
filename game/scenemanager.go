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
	transitioning bool
}

func NewSceneManager() *SceneManager {
	return &SceneManager{}
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

	// Draw the previous screen in case we change screens in the middle of update
	// If there is a change in the middle of update then the new screen's update will not have run
	// and a bunch of stuff won't have been initialized
	if s.transitioning == true {
		s.previous.Draw(state, screen)
		s.transitioning = false
	} else {
		s.current.Draw(state, screen)
	}
}

func (s *SceneManager) GoTo(scene Scene) {
	// If this isn't the first scene make sure we set the previous
	if s.current != nil {
		s.transitioning = true
		s.previous = s.current
	}

	s.current = scene
}
