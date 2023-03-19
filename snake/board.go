package snake

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Board struct {
	rows      int    // 行
	cols      int    // 列
	food      *Food  // 食物
	snake     *Snake // 🐍
	scores    int    // 分数
	gameStart bool   // 游戏开始
	gameOver  bool   // 游戏结束
	timer     time.Time
}

// NewBoard 创建一个新的 Board
func NewBoard(rows, cols int) *Board {
	b := &Board{
		rows:      rows,
		cols:      cols,
		timer:     time.Now(),
		gameStart: false,
		gameOver:  false,
	}
	b.snake = NewSnake([]Coord{{0, 0}, {1, 0}, {2, 0}, {3, 0}}, ebiten.KeyArrowRight)
	b.placeFood() // 放食物
	return b
}

// Update 更新Board
func (b *Board) Update(i *Input) error {
	// 游戏开始
	if ok := i.isPressSpace(); ok {
		b.gameStart = true
	}
	// 重新开始
	if ok := i.isPressR(); ok {
		b.gameStart = false
		b.gameOver = false
		b.snake = NewSnake([]Coord{{0, 0}, {1, 0}, {2, 0}, {3, 0}}, ebiten.KeyArrowRight)
	}
	// 游戏结束
	if b.gameOver {
		return nil
	}
	// 改变方向
	if newDir, ok := i.Dir(); ok {
		b.snake.ChangeDirection(newDir)
	}

	// 移动🐍身
	interval := time.Millisecond * 200
	if b.gameStart && time.Since(b.timer) >= interval {
		if err := b.moveSnake(); err != nil {
			return err
		}
		b.timer = time.Now()
	}
	return nil
}

// DisplayStartScreen 在screen上展示游戏开始界面
func (b *Board) DisplayStartScreen(screen *ebiten.Image, face font.Face) {
	text.Draw(
		screen,
		"Press the space key to start the game!\n",
		face,
		ScreenWidth/2-fontSize*8, ScreenHeight/2,
		color.Black,
	)
}

// DisplayScore 在screen上显示分数
func (b *Board) DisplayScore(screen *ebiten.Image, score int, face font.Face) {
	text.Draw(
		screen,
		fmt.Sprintf("Score: %d", score),
		face,
		0, fontSize,
		color.Black,
	)
}

// DisplayOverScreen 在screen上展示游结束界面
func (b *Board) DisplayOverScreen(screen *ebiten.Image, score int, face font.Face) {
	text.Draw(
		screen,
		fmt.Sprintf("Game Over. Score: %d\n", score)+
			"Press R to restart the game.\n",
		face,
		ScreenWidth/2-fontSize*8, ScreenHeight/2,
		color.Black,
	)
}

// DisplaySnake 画🐍
func (b *Board) DisplaySnake(screen *ebiten.Image) {
	snakeColor := snakeBodyColor
	for i, p := range b.snake.body {
		if i == len(b.snake.body)-1 {
			snakeColor = snakeHeadColor
		}
		ebitenutil.DrawRect(screen, float64(p.x*coordWidth)+float64(coordWidth*1/20), float64(p.y*coordHeight)+float64(coordHeight*1/20), float64(coordWidth)*9/10, float64(coordHeight)*9/10, snakeColor)
	}
}

// DisplayFood 画食物
func (b *Board) DisplayFood(screen *ebiten.Image) {
	foodImg := b.food.image
	op := &ebiten.DrawImageOptions{}
	sx, sy := foodImg.Size()
	propx := float64(coordWidth) / float64(sy)
	propy := float64(coordHeight) / float64(sx)
	op.GeoM.Scale(propx, propy)
	op.GeoM.Translate(float64(b.food.x*coordWidth), float64(b.food.y*coordHeight))
	screen.DrawImage(foodImg, op)
}

// placeFood 放置食物
func (b *Board) placeFood() {
	var x, y int
	for {
		x = rand.Intn(b.cols)
		y = rand.Intn(b.rows)
		on := false // 食物是否在🐍上
		for _, v := range b.snake.body {
			if x == v.x && y == v.y {
				on = true
			}
		}
		if !on && !b.snake.HeadHits(x, y) {
			break
		}
	}
	b.food = NewFood(x, y)
}

// moveSnake 移动🐍
func (b *Board) moveSnake() error {
	b.snake.Move() // 移动
	if b.isTouchTheWall() || b.snake.HeadHitsBody() {
		b.gameOver = true
		return nil
	}
	if b.snake.HeadHits(b.food.x, b.food.y) {
		b.snake.justEat = true // 是否吃到食物
		b.placeFood()          // 放食物
		b.scores++             // 分数
	}
	return nil
}

// isTouchTheWall是否碰撞墙壁
func (b *Board) isTouchTheWall() bool {
	head := b.snake.Head()
	return head.x < 0 || head.y < 0 || head.x > b.cols-1 || head.y > b.rows-1
}
