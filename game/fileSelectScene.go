package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type FileSelectScene struct {
	background      *Background
	initialized     bool
	currentElement  int
	previousElement int
	backButton      *BackButton
	optionsButton   *OptionsButton
	deleteButton    *DeleteButton
}

func (s *FileSelectScene) initialize(state *GameState) {
	s.backButton = NewBackButton(
		state.sceneManager.previous,
		func() (int, int) { width, height := Resolution(); return width * 9 / 10, height * 9 / 10 })

	s.currentElement = 0
	s.previousElement = 1 // Set this to 1 so that there is no conflict with currentElement
	s.optionsButton = NewOptionsButton(func() (int, int) { width, height := Resolution(); return width * 1 / 10, height * 9 / 10 })

	s.initialized = true
}

func (s *FileSelectScene) elementCount() int {
	// Back Button, Save Button, Settings
	return 2
}

// Highlight the currently selected Option
func (s *FileSelectScene) highlightOption() {
	// Highlight of backbutton
	if s.currentElement == 0 {
		s.backButton.Highlight(true)
	} else {
		s.backButton.Highlight(false)
	}

	// Highlight of saveButton
	if s.currentElement == 1 {
		s.optionsButton.Highlight(true)
	} else {
		s.optionsButton.Highlight(false)
	}
}

func (s *FileSelectScene) Update(state *GameState) error {
	if s.initialized == false {
		s.initialize(state)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		s.previousElement = s.currentElement
		s.currentElement -= 1
		if s.currentElement < 0 {
			s.currentElement = s.elementCount() - 1
		}
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		s.previousElement = s.currentElement
		s.currentElement += 1
		if s.currentElement >= s.elementCount() {
			s.currentElement = 0
		}
		return nil
	}

	s.optionsButton.Update(state)
	s.backButton.Update(state)

	// Highlight new and unhighlight old
	s.highlightOption()

	return nil
}

func (s *FileSelectScene) Draw(screen *ebiten.Image) {
	gw, gh := Resolution()

	drawbackground(screen, gw, gh)
	drawText(screen, gw/10, gh/10, "File Select", color.Black)

	s.optionsButton.Draw(screen)
	s.backButton.Draw(screen)
}

type DeleteButton struct {
}
