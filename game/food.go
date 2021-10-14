package game

import (
	"math/rand"
	"time"
)

func init()  {
	rand.Seed(time.Now().UnixNano())
}

type food struct {
	pos coord
}

func createFood(xLimit, yLimit int) food {
	xLimit--
	yLimit--
	x := rand.Intn(xLimit - 1) + 1
	y := rand.Intn(yLimit - 1) + 1

	return food{
		pos: coord{
			x: x,
			y: y,
		},
	}
}
