package snake

type Food struct {
	x int
	y int
}

func NewFood(x, y int) *Food {
	return &Food{
		x: x,
		y: y,
	}
}
