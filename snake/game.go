package snake

import (
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 600
	ScreenHeight = 420
	boardCols    = 20
	boardRows    = 14
	coordWidth   = ScreenWidth / boardCols       // 每个小方块的宽度
	coordHeight  = ScreenHeight / boardRows      // 每个小方块的高度
	fontSize     = 25                            // 字体大小
	name         = "assets/background music.mp3" // 背景音乐
	sampleRate   = 48000                         // 码率
)

var (
	backgroundColor = color.RGBA{83, 175, 74, 255}
	snakeHeadColor  = color.RGBA{21, 21, 43, 255} // 🐍头颜色
	snakeBodyColor  = color.RGBA{21, 21, 43, 255} // 🐍身颜色
)

type Game struct {
	input            *Input    // 输入
	board            *Board    // 背板
	score            *Score    // 分数
	isGameOver       bool      // 游戏是否结束界面
	isGameStart      bool      // 游戏是否开始界面
	isGameInProgress bool      // 游戏是否正在进行
	startTime        time.Time // 游戏开始时间
}

func NewGame() *Game {
	return &Game{
		input:       NewInput(),
		score:       NewScore(),
		board:       NewBoard(boardRows, boardCols),
		isGameStart: true,
	}
}

func (g *Game) Update() error {
	var err error
	// 重新开始
	if ok := g.isGameOver && g.input.isPressR(); ok {
		g.board = NewBoard(boardRows, boardCols)
		g.isGameStart = true
		g.isGameInProgress = false
		g.isGameOver = false
	}
	// 游戏结束
	if g.isGameOver {
		return nil
	}
	// 更新状态 isGameOver, isGameStart, isGameInProgress
	// 游戏开始
	if ok := g.isGameStart && g.input.isPressSpace(); ok {
		if g.score.score != 0 {
			g.score.Save()
		}
		g.score.score = 0
		g.isGameStart = false
		g.isGameInProgress = true
		g.isGameOver = false
		g.startTime = time.Now()
	}

	if g.isGameInProgress { // 更新游戏进行时
		// 改变方向
		if newDir, ok := g.input.Dir(); ok {
			g.board.snake.ChangeDirection(newDir)
		}
		// 移动🐍身
		interval := time.Millisecond * 200
		if time.Since(g.board.timer) >= interval {
			if err := g.board.MoveSnake(g); err != nil {
				return err
			}
			g.board.timer = time.Now()
		}
	}
	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor) // 填充背景
	g.board.DrawGrid(screen)
	// face, err := loadGoregularFont(fontSize) // Goregular字体
	face, err := loadLocalFont(fontSize)
	if err != nil {
		log.Fatal(err)
	}
	if g.isGameStart { // 游戏开始界面
		g.board.DisplayStartScreen(screen, face)
	}
	if g.isGameOver { // 游戏结束 显示分数
		g.board.DisplayOverScreen(screen, g.score.score, g.score.HighestScore(), face)
	}
	if g.isGameInProgress {
		g.board.DisplaySnake(screen)                        // 画🐍身
		g.board.DisplayFoods(screen)                        // 画食物
		g.board.DisplayScore(screen, g.score.score, face)   // 实时分数
		g.board.DisplaySpentTime(screen, g.startTime, face) // 用时
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
