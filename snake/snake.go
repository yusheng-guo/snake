package snake

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Snake struct {
	body      []Coord    // 身体
	direction ebiten.Key //方向
	justEat   bool
	sounds    map[string]*Sound // 声音
}

// NewSnake 创建🐍
func NewSnake(body []Coord, direction ebiten.Key) *Snake {
	var err error
	m := &Snake{
		body:      body,
		direction: direction,
	}
	m.sounds, err = LoadSounds()
	if err != nil {
		log.Fatal(err)
	}
	return m
}

// Head 🐍头
func (s *Snake) Head() Coord {
	return s.body[len(s.body)-1]
}

// ChangeDirection 改变🐍的方向
// 不允许将方向修改为原方向的相反方向
func (s *Snake) ChangeDirection(newDir ebiten.Key) {
	opposites := map[ebiten.Key]ebiten.Key{
		ebiten.KeyArrowUp:    ebiten.KeyArrowDown,
		ebiten.KeyArrowRight: ebiten.KeyArrowLeft,
		ebiten.KeyArrowDown:  ebiten.KeyArrowUp,
		ebiten.KeyArrowLeft:  ebiten.KeyArrowRight,
	}
	if d, ok := opposites[newDir]; ok && d != s.direction {
		s.direction = newDir
	}
}

// 碰撞检测
// HeadHits 检测🐍头是否在(x, y)
// 是否吃到食物
func (s *Snake) HeadHits(pos Coord) bool {
	head := s.Head()
	return head.x == pos.x && head.y == pos.y
}

// HeadHits 检测🐍头是否碰撞🐍身
func (s *Snake) HeadHitsBody() bool {
	head := s.Head()
	bodyWithoutHead := s.body[:len(s.body)-1]
	for _, b := range bodyWithoutHead {
		if b.x == head.x && b.y == head.y {
			s.playSound("over")
			return true
		}
	}
	return false
}

// Move 🐍移动
func (s *Snake) Move() {
	head := s.Head() // 🐍头
	newHead := Coord{
		x: head.x,
		y: head.y,
	}
	switch s.direction {
	case ebiten.KeyArrowDown:
		newHead.y++
	case ebiten.KeyArrowUp:
		newHead.y--
	case ebiten.KeyArrowLeft:
		newHead.x--
	case ebiten.KeyArrowRight:
		newHead.x++
	}
	if s.justEat {
		s.body = append(s.body, newHead)
		s.justEat = false
	} else {
		s.body = append(s.body[1:], newHead)
	}
}

func (s *Snake) playSound(name string) {
	if s, ok := s.sounds[name]; ok {
		if err := s.Play(); err != nil {
			log.Fatal(err)
		}
	}
}
