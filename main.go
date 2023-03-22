package main

import (
	"bytes"
	_ "embed"
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yushengguo557/snake/snake"
)

//go:embed .\assets\icon.png
var icon []byte

func main() {
	os.Setenv("CGO_ENABLED", "1")
	game := snake.NewGame()
	ebiten.SetWindowSize(snake.ScreenWidth, snake.ScreenHeight)
	ebiten.SetWindowTitle("Snake")
	reader := bytes.NewReader(icon)
	img, _, err := ebitenutil.NewImageFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowIcon([]image.Image{img})
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
