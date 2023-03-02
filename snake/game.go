package snake

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 500 // 屏幕宽度
	ScreenHeight = 500 // 屏幕高度
	boardRows    = 20  // 板 行数
	boardCols    = 20  // 板 列数
)

var (
	backgroundColor = color.RGBA{7, 193, 96, 100}   // 背景颜色
	snakeColor      = color.RGBA{200, 50, 150, 150} // 🐍颜色
	foodColor       = color.RGBA{200, 200, 50, 150} // 食物颜色
)

type Game struct {
	input *Input // 输入
	board *Board // 板子
}

func NewGame() *Game {
	return &Game{
		input: NewInput(),
		board: NewBoard(boardRows, boardCols),
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	return g.board.Update(g.input)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	if g.board.gameOver {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Game Over. Score: %d", g.board.points))
	} else {
		width := ScreenHeight / boardRows

		for _, p := range g.board.snake.body {
			ebitenutil.DrawRect(screen, float64(p.y*width), float64(p.x*width), float64(width), float64(width), snakeColor)
		}
		if g.board.food != nil {
			ebitenutil.DrawRect(screen, float64(g.board.food.y*width), float64(g.board.food.x*width), float64(width), float64(width), foodColor)
		}
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.board.points))
	}
}
