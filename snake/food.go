package snake

import (
	"log"
	"math/rand"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Food struct {
	x     int
	y     int
	image *ebiten.Image
}

func NewFood(x, y int) *Food {
	return &Food{
		x:     x,
		y:     y,
		image: getRandFoodImage(),
	}
}

func getRandFoodImage() *ebiten.Image {
	dirs, err := os.ReadDir("assets/foods/")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := ebitenutil.NewImageFromFile("assets/foods/" + dirs[rand.Intn(len(dirs))].Name())
	if err != nil {
		log.Fatal(err)
	}
	return img
}
