package game

import (
	"math/rand"
	"time"
)

func init()  {
	rand.Seed(time.Now().UnixNano())
}

type Food struct {
	Pos coord
}

func CreateFood(xLimit, yLimit int) Food {
	xLimit--
	yLimit--
	//if xLimit == 0 {
	//	xLimit = 1
	//}
	//if yLimit == 0 {
	//	yLimit = 1
	//}
	x := rand.Intn(xLimit - 1) + 1
	y := rand.Intn(yLimit - 1) + 1

	return Food{
		Pos: coord{
			x: x,
			y: y,
		},
	}
}
