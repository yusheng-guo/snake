package snake

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 600
	ScreenHeight = 420
	boardCols    = 20
	boardRows    = 14
	coordWidth   = ScreenWidth / boardCols                        // 每个小方块的宽度
	coordHeight  = ScreenHeight / boardRows                       // 每个小方块的高度
	fontSize     = 20                                             // 字体大小
	name         = "assets/Lewis Capaldi - Someone You Loved.mp3" // 背景音乐
)

var (
	backgroundColor = color.RGBA{128, 128, 128, 255}
	snakeHeadColor  = color.RGBA{255, 99, 71, 255}   // 🐍头颜色
	snakeBodyColor  = color.RGBA{149, 236, 105, 255} // 🐍身颜色
)

type Game struct {
	input *Input // 输入
	board *Board // 背板
	music *Music // 音乐
}

func NewGame() *Game {
	return &Game{
		input: NewInput(),
		board: NewBoard(boardRows, boardCols),
		music: NewMusic(name),
	}
}

func (g *Game) Update() error {
	g.music.player.Play()
	return g.board.Update(g.input)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor) // 填充背景
	// face, err := loadGoregularFont(fontSize) // Goregular字体
	face, err := loadLocalFont("assets/Comic Sans MS.ttf", fontSize)
	if err != nil {
		log.Fatal(err)
	}
	if !g.board.gameStart { // 游戏开始界面
		g.board.DisplayStartScreen(screen, face)
	} else if g.board.gameOver { // 游戏结束 显示分数
		g.board.DisplayOverScreen(screen, g.board.scores, face)
	} else {
		g.board.DisplaySnake(screen)                       // 画🐍身
		g.board.DisplayFood(screen)                        // 画食物
		g.board.DisplayScore(screen, g.board.scores, face) // 实时分数
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
