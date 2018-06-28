package game
import (
  "image/color"
)

type Text struct {
  content string
  opts *TextOpts
}

type TextOpts struct {
  clr color.Color
  flashing bool
  floating bool
}

func NewText(content string, options *TextOpts) *Text {
  return &Text{content: content, opts: options}
}

func (t *Text) Update() {

}

func (t *Text) Draw() {

}
