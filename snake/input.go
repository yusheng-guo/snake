package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct{}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) Dir() (ebiten.Key, bool) {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		return ebiten.KeyArrowUp, true
	}
	if inpututil.IsKeyJustPressed((ebiten.KeyArrowDown)) {
		return ebiten.KeyArrowDown, true
	}
	if inpututil.IsKeyJustPressed((ebiten.KeyArrowLeft)) {
		return ebiten.KeyArrowLeft, true
	}
	if inpututil.IsKeyJustPressed((ebiten.KeyArrowRight)) {
		return ebiten.KeyArrowRight, true
	}
	return 0, false
}

func (i *Input) isPressSpace() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeySpace)
}

func (i *Input) isPressR() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyR)
}
