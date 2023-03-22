package snake

import (
	"bytes"
	_ "embed"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed ..\assets\foods\cake.png
var cake []byte

//go:embed ..\assets\foods\apple.png
var apple []byte

//go:embed ..\assets\foods\cookieChocolate.png
var cookieChocolate []byte

//go:embed ..\assets\foods\fish.png
var fish []byte

//go:embed ..\assets\foods\loaf.png
var loaf []byte

//go:embed ..\assets\foods\turkey.png
var turkey []byte

var foodFiles = [][]byte{
	cake,
	apple,
	cookieChocolate,
	fish,
	loaf,
	turkey,
}

type Food struct {
	position Coord
	image    *ebiten.Image
}

type Foods struct {
	foods  []*Food
	images []*ebiten.Image
}

func NewFood(x, y int, image *ebiten.Image) *Food {
	return &Food{
		position: Coord{x: x, y: y},
		image:    image,
	}
}

func NewFoods() *Foods {
	return &Foods{
		images: loadFoodsPictures(),
	}
}

// getRandFoodImage 获取随机食物图片
// func getRandFoodImage() *ebiten.Image {
// 	dirs, err := os.ReadDir("assets/foods/")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	img, _, err := ebitenutil.NewImageFromFile("assets/foods/" + dirs[rand.Intn(len(dirs))].Name())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return img
// }

func loadFoodsPictures() []*ebiten.Image {
	var pictures []*ebiten.Image
	for _, f := range foodFiles {
		reader := bytes.NewReader(f)
		img, _, err := ebitenutil.NewImageFromReader(reader)
		if err != nil {
			log.Fatal(err)
		}
		pictures = append(pictures, img)

	}

	return pictures
}

// getRandFoodImage 随机获取一张图片
func (foods *Foods) getRandFoodImage() *ebiten.Image {
	return foods.images[rand.Intn(len(foods.images))]
}
