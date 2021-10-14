package game

import (
	"fmt"
	"github.com/Kolakanmi/snake/utils"
	"os"
	"os/exec"
	"time"
)

type Board struct {
	width int
	height int
	stage [][]int
	snake *Snake
	food Food
	score int
	gameOver bool
	stop bool
	dir Direction
	speed int
	currentRound int64
}

func (b *Board) Run() {
	go func() {
		b.Input()
	}()
	for !b.stop {
		for !b.gameOver {
			b.DisplayStage()
		}
	}

}

func (b *Board) DisplayStage2()  {
	for i, x := range b.stage{
		for j := range x {
			if b.stage[i][j] == 1 && j == len(x) - 1{
				fmt.Println("*")
			} else if b.stage[i][j] == 1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
	}
}

func (b *Board) SetIndexValue(x, y, val int) {
	b.stage[x][y] = val
}

func (b *Board) clearFood() {
	x, y := b.food.Pos.Get()
	b.SetIndexValue(x, y, 0)
}

func (b *Board) setSnake(dir Direction) {
	x, y := b.snake.Head.Get()
	x0, y0 := x, y
	b.SetIndexValue(x, y, 0)
	for i := 0; i < len(b.snake.Tail); i++ {
		txn, tyn := b.snake.Tail[i].Get()
		b.SetIndexValue(txn, tyn, 0)
	}
	dirTemp := dir
	switch dir {
	case left:
		y--
	case right:
		y++
	case up:
		x--
	case down:
		x++
	}
	//b.dir = stop

	if x <= 0 {
		x = len(b.stage) - 2
	}
	if x >= len(b.stage) - 1 {
		x = 1
	}
	if y <= 0 {
		y = len(b.stage[0]) - 2
	}
	if y >= len(b.stage[0]) - 1{
		y = 1
	}
	b.snake.Head.Set(x, y)
	b.SetIndexValue(x, y, 2)

	xp, yp := 0, 0
	for i := 0; i < len(b.snake.Tail); i++ {
		xc, yc := b.snake.Tail[i].Get()
		if i == 0 {
			b.snake.Tail[i].Set(x0, y0)
			b.SetIndexValue(x0, y0, 2)
		} else {
			b.snake.Tail[i].Set(xp, yp)
			b.SetIndexValue(xp, yp, 2)
		}
		xp, yp = xc, yc
	}

	for _, t := range b.snake.Tail {
		if x == t.x && y == t.y {
			b.gameOver = true
		}
	}

	xFood, yFood := b.food.Pos.Get()

	if x == xFood && y == yFood {
		//b.food = CreateFood(b.width, b.height)
		//
		//xFood, yFood := b.food.Pos.Get()
		//b.stage[xFood][yFood] = 3
		b.newFood()
		b.score += 1
		b.snake.AddTail(dirTemp)

		xl, yl := b.snake.GetLastTail()
		b.SetIndexValue(xl, yl, 2)
	}
}

func (b *Board) newFood() {
	b.food = CreateFood(b.width, b.height)
	xFood, yFood := b.food.Pos.Get()
	retry := false
	for _, v := range b.snake.Tail {
		xt, yt := v.Get()
		if xt == xFood && yt == yFood {
			retry = true
			break
		}
	}
	if !retry {
		b.SetIndexValue(xFood, yFood, 3)
	} else  {
		b.newFood()
	}
}

func (b *Board) Input() {
	ch := make(chan string)
	cmd1 := exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1")
	cmd1.Run()
	cmd2 := exec.Command("stty", "-F", "/dev/tty", "-echo")
	cmd2.Run()
	go func(ch chan string) {
		var bb []byte = make([]byte, 1)
		for  {
			os.Stdin.Read(bb)
			ch <- string(bb)
		}
	}(ch)

	for  {
		select {
		case stdin, _ := <-ch:
			switch stdin {
			case "a":
				if b.dir != right {
					b.dir = left
				}
			case "d":
				if b.dir != left {
					b.dir = right
				}
			case "w":
				if b.dir != down {
					b.dir = up
				}
			case "s":
				if b.dir != up {
					b.dir = down
				}
			case "x":
				b.gameOver = true
				b.stop = true
			case "r":
				b.Restart()
			}
		}
	}
}

func (b *Board) Logic() {
	if b.dir != stop {
		b.setSnake(b.dir)
	}
}

func (b *Board) DisplayStage()  {
	utils.ClearTerminal()
	b.Logic()

	utils.ClearTerminal()
	for i, x := range b.stage{
		for j := range x {
			if b.stage[i][j] == 1 && j == len(x) - 1{
				fmt.Println("*")
			} else if b.stage[i][j] == 1 && (i == 0 || i == len(x) - 1 || j == 0) {
				fmt.Print("*")
			} else if b.stage[i][j] == 1 {
				fmt.Print("*")
			} else if b.stage[i][j] == 2 {
				fmt.Print("o")
			} else if b.stage[i][j] == 3 {
				fmt.Print("x")
			} else {
				fmt.Print(" ")
			}
		}
	}

	fmt.Println()
	fmt.Println()
	if !b.gameOver {
		fmt.Printf("%d%+v", b.score, b.snake.Head)
	} else {
		fmt.Println("GAME OVER.")
		fmt.Println("PRESS 'r' to Restart")
		fmt.Println("PRESS 'x' to End")
	}

	time.Sleep(100* time.Millisecond)
}

func CreateBoard(width int) *Board {
	b := &Board{
		width: width,
		height: width * 3,
		dir: up,
		stop: false,
	}
	b.SetStage()
	return b
}

func (b *Board) Restart() {
	b.SetStage()
	b.score = 0
	b.stop = false
	b.gameOver = false
}

func (b *Board) SetStage() {
	b.stage = make([][]int, b.width)
	for i := range b.stage {
		b.stage[i] = make([]int, b.height)
	}

	b.snake = &Snake{
		length: 1,
		Head: &coord{
			x: b.width/2,
			y: b.height/2,
		},
	}

	b.food = CreateFood(b.width, b.height)
	//food := CreateFood(b.width, b.height)
	//b.food = food

	b.SetBorder()
	xFood, yFood := b.food.Pos.Get()
	b.stage[xFood][yFood] = 3
}

func (b *Board) SetBorder() {
	for i, x := range b.stage{
		for j := range x {
			if i == 0 || i == len(b.stage) -1 {
				b.stage[i][j] = 1
				continue
			} else if j == 0 || j == len(x) -1 {
				b.stage[i][j] = 1
				continue
			}
			b.stage[i][j] = 0
		}
	}
}
