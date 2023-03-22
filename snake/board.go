package snake

import (
	"fmt"
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/exp/slices"
	"golang.org/x/image/font"
)

type Board struct {
	rows  int       // 行
	cols  int       // 列
	foods *Foods    // 食物
	snake *Snake    // 🐍
	timer time.Time // 控制🐍移动速度
}

// NewBoard 创建一个新的 Board
func NewBoard(rows, cols int) *Board {
	b := &Board{
		rows:  rows,
		cols:  cols,
		timer: time.Now(),
		foods: NewFoods(),
	}
	b.snake = NewSnake([]Coord{{-3, 0}, {-2, 0}, {-1, 0}, {0, 0}}, ebiten.KeyArrowRight)
	b.placeFoods(5) // 放食物
	return b
}

// DrawGrid 画网格
func (b *Board) DrawGrid(screen *ebiten.Image) {
	// 画点
	// for x := 1; x < b.cols; x++ {
	// 	for y := 1; y <= b.rows; y++ {
	// 		ebitenutil.DrawRect(screen, float64(x*coordWidth), float64(y*coordHeight), float64(2), float64(2), color.Black)
	// 	}
	// }
	// 画线
	for x := 1; x < b.cols; x++ {
		ebitenutil.DrawLine(screen, float64(x*coordWidth), float64(0), float64(x*coordWidth), float64(ScreenHeight), color.RGBA{128, 128, 128, 255})
	}
	for y := 1; y < b.cols; y++ {
		ebitenutil.DrawLine(screen, float64(0), float64(y*coordHeight), float64(ScreenWidth), float64(y*coordHeight), color.RGBA{128, 128, 128, 255})
	}
}

// DisplayStartScreen 在screen上展示游戏开始界面
func (b *Board) DisplayStartScreen(screen *ebiten.Image, face font.Face) {
	message := "Press the \"space\" key to start the game!\n"
	size := text.BoundString(face, message)
	messageWidth, messageHeight := size.Max.X-size.Min.X, size.Max.Y-size.Min.Y
	text.Draw(
		screen,
		message,
		face,
		(ScreenWidth-messageWidth)/2, (ScreenHeight-messageHeight)/2,
		color.Black,
	)
}

// DisplayScore 在screen上显示分数
func (b *Board) DisplayScore(screen *ebiten.Image, score int, face font.Face) {
	message := fmt.Sprintf("Score: %d", score)
	text.Draw(
		screen,
		message,
		face,
		ScreenWidth-150, 30,
		color.RGBA{255, 0, 0, 255},
	)
}

// DisplaySpentTime 在screen上显示用时
func (b *Board) DisplaySpentTime(screen *ebiten.Image, startTime time.Time, face font.Face) {
	message := fmt.Sprintf("Spent: %.0fs", time.Since(startTime).Seconds())
	text.Draw(
		screen,
		message,
		face,
		ScreenWidth-150, 60,
		color.RGBA{255, 0, 0, 255},
	)
}

// DisplayOverScreen 在screen上展示游结束界面
func (b *Board) DisplayOverScreen(screen *ebiten.Image, score int, highestScore int, face font.Face) {
	message := "Game Over.\n" + fmt.Sprintf("Score: %d\n", score) + "Max Score: " + strconv.Itoa(highestScore) + "\n" + "Press R to restart the game.\n"
	size := text.BoundString(face, message)
	messageWidth, messageHeight := size.Max.X-size.Min.X, size.Max.Y-size.Min.Y
	text.Draw(
		screen,
		message,
		face,
		(ScreenWidth-messageWidth)/2, (ScreenHeight-messageHeight)/2,
		color.RGBA{220, 20, 60, 255},
	)
}

// DisplaySnake 画🐍
func (b *Board) DisplaySnake(screen *ebiten.Image) {
	snakeColor := snakeBodyColor
	for i, p := range b.snake.body {
		if i == len(b.snake.body)-1 {
			snakeColor = snakeHeadColor
		}
		// ebitenutil.DrawRect(screen, float64(p.x*coordWidth)+float64(coordWidth*1/20), float64(p.y*coordHeight)+float64(coordHeight*1/20), float64(coordWidth)*9/10, float64(coordHeight)*9/10, snakeColor)
		ebitenutil.DrawCircle(
			screen,
			float64(p.x*coordWidth)+float64(coordWidth/2), float64(p.y*coordHeight)+float64(coordHeight/2),
			float64(coordWidth/2),
			snakeColor,
		)
	}
}

// DisplayFood 画食物
func (b *Board) DisplayFoods(screen *ebiten.Image) {
	for _, f := range b.foods.foods {
		b.displayFood(screen, f)
	}
}

func (b *Board) displayFood(screen *ebiten.Image, food *Food) {
	foodImg := food.image
	op := &ebiten.DrawImageOptions{}
	sx, sy := foodImg.Size()
	propx := float64(coordWidth) / float64(sy)
	propy := float64(coordHeight) / float64(sx)
	op.GeoM.Scale(propx, propy)
	op.GeoM.Translate(float64(food.position.x*coordWidth), float64(food.position.y*coordHeight))
	screen.DrawImage(foodImg, op)
}

// placeFoods 放置n个食物
func (b *Board) placeFoods(n int) {
	var x, y int
	for i := 0; i < n; i++ {
		for {
			x = rand.Intn(b.cols)
			y = rand.Intn(b.rows)
			on := false // 食物是否在🐍上
			for _, v := range b.snake.body {
				if x == v.x && y == v.y {
					on = true
				}
			}
			if !on && !b.snake.HeadHits(Coord{x, y}) {
				break
			}
		}
		b.foods.foods = append(b.foods.foods, NewFood(x, y, b.foods.getRandFoodImage()))
	}
}

// MoveSnake 移动🐍
func (b *Board) MoveSnake(g *Game) error {
	b.snake.Move()                                    // 移动
	if b.isTouchTheWall() || b.snake.HeadHitsBody() { // 游戏结束
		b.snake.playSound("over")
		// 更新游戏状态
		g.isGameOver = true
		g.isGameInProgress = false
		g.isGameStart = false
		return nil
	}
	for _, f := range b.foods.foods {
		if b.snake.HeadHits(Coord{f.position.x, f.position.y}) {
			if score, ok := b.snake.sounds["score"]; ok { // 吃到食物音效
				score.Play()
			}
			// 更新食物链
			index := slices.IndexFunc(b.foods.foods, func(f *Food) bool {
				return f.position.x == b.snake.Head().x && f.position.y == b.snake.Head().y
			})
			if index == len(b.foods.foods)-1 {
				b.foods.foods = b.foods.foods[:index]
			} else {
				b.foods.foods = append(b.foods.foods[:index], b.foods.foods[index+1:]...)
			}
			b.snake.justEat = true // 是否吃到食物
			b.placeFoods(1)        // 放食物
			g.score.score++        // 分数
		}
	}
	return nil
}

// isTouchTheWall是否碰撞墙壁
func (b *Board) isTouchTheWall() bool {
	head := b.snake.Head()
	return head.x < 0 || head.y < 0 || head.x > b.cols-1 || head.y > b.rows-1
}
