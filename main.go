package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yushengguo557/snake/snake"
)

func main() {
	game := snake.NewGame()
	ebiten.RunGame(game)
}
