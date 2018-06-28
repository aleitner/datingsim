package game

import (
	"image"
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/nfnt/resize"
)

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

	return &Background{img: img}, nil
}

func (b *Background) Update(state *GameState) {
}

func (b *Background) Draw(screen *ebiten.Image) {
	if b.img == nil {
		return
	}

	gw, gh := Resolution()

	// Resize the image to match the resolution
	resizedIMG := resize.Resize(uint(gw), 0, b.img, resize.Lanczos3)

	// Convert the imageto ebiten format
	img, err := ebiten.NewImageFromImage(resizedIMG, ebiten.FilterDefault)
	if err != nil {
		return
	}

	_, ih := img.Size()
	op := &ebiten.DrawImageOptions{}

	// translate the image to the proper y coordinate if the image is shorter than the height
	if heightDiff := gh - ih; heightDiff > 0 {
		op.GeoM.Translate(0, float64(heightDiff/2))
	}

	screen.DrawImage(img, op)
}
