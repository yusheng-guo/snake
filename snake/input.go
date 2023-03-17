package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct{}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) Dir(ebiten.Key) (ebiten.Key, bool) {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		return ebiten.KeyUp, true
	}
	if inpututil.IsKeyJustPressed((ebiten.KeyDown)) {
		return ebiten.KeyDown, true
	}
	if inpututil.IsKeyJustPressed((ebiten.KeyLeft)) {
		return ebiten.KeyLeft, true
	}
	if inpututil.IsKeyJustPressed((ebiten.KeyRight)) {
		return ebiten.KeyRight, true
	}
	return 0, false
}
