package game

import (
	"fmt"
	"github.com/Kolakanmi/snake/utils"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type board struct {
	width        int
	height       int
	stage        [][]int
	snake        *snake
	food         food
	score        int
	gameOver     bool
	stop         bool
	dir          direction
	speed        int
	currentRound int64
}

func Run() {
	size := 30
	var input string
	valid := false
	for !valid {
		fmt.Println()
		fmt.Printf("%s\n\n","Board width will be three times more than height")
		fmt.Print("Enter board height: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Enter a number.")
		} else {
			num, err := strconv.Atoi(input)
			if err != nil {
				fmt.Printf("%s\n\n", "ENTER A NUMBER.")
			} else {
				if num < 10 {
					fmt.Printf("%s\n\n", "MINIMUM HEIGHT IS 10")
				} else {
					size = num
					valid = true
				}
			}
		}
	}

	b := createBoard(size)
	go func() {
		b.input()
	}()
	for !b.stop {
		for !b.gameOver {
			b.displayStage()
		}
	}

}

func (b *board) setIndexValue(x, y, val int) {
	b.stage[y][x] = val
}

func (b *board) clearFood() {
	x, y := b.food.pos.get()
	b.setIndexValue(x, y, 0)
}

func (b *board) setSnake(dir direction) {
	x, y := b.snake.Head.get()
	x0, y0 := x, y
	b.setIndexValue(x, y, 0)
	for i := 0; i < len(b.snake.Tail); i++ {
		txn, tyn := b.snake.Tail[i].get()
		b.setIndexValue(txn, tyn, 0)
	}
	dirTemp := dir
	b.currentRound += 1
	switch dir {
	case left:
		x--
	case right:
		x++
	case up:
		y--
	case down:
		y++
	}
	//b.dir = stop

	if y <= 0 {
		y = len(b.stage) - 2
	}
	if y >= len(b.stage) - 1 {
		y = 1
	}
	if x <= 0 {
		x = len(b.stage[0]) - 2
	}
	if x >= len(b.stage[0]) - 1{
		x = 1
	}
	b.snake.Head.set(x, y)
	b.setIndexValue(x, y, 4)

	xp, yp := 0, 0
	for i := 0; i < len(b.snake.Tail); i++ {
		xc, yc := b.snake.Tail[i].get()
		if i == 0 {
			b.snake.Tail[i].set(x0, y0)
			b.setIndexValue(x0, y0, 2)
		} else {
			b.snake.Tail[i].set(xp, yp)
			b.setIndexValue(xp, yp, 2)
		}
		xp, yp = xc, yc
	}

	xFood, yFood := b.food.pos.get()

	if x == xFood && y == yFood {
		b.newFood()
		b.score += 1
		b.snake.addTail(dirTemp)

		xl, yl := b.snake.getLastTail()
		b.setIndexValue(xl, yl, 2)
	}
	for _, t := range b.snake.Tail {
		if x == t.x && y == t.y {
			b.gameOver = true
		}
	}
}

func (b *board) newFood() {
	b.food = createFood(b.width, b.height)
	xFood, yFood := b.food.pos.get()
	retry := false
	for _, v := range b.snake.Tail {
		xt, yt := v.get()
		if xt == xFood && yt == yFood {
			retry = true
			break
		}
	}
	if !retry {
		b.setIndexValue(xFood, yFood, 3)
	} else  {
		b.newFood()
	}
}

func (b *board) input() {
	ch := make(chan string)
	cmd1 := exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1")
	err := cmd1.Run()
	if err != nil {
		panic("Could not run command to read input without pressing 'Enter key'.")
	}
	cmd2 := exec.Command("stty", "-F", "/dev/tty", "-echo")
	err = cmd2.Run()
	if err != nil {
		panic("Could not run command to read input without pressing 'Enter key'.")
	}
	go func(ch chan string) {
		var bb = make([]byte, 1)
		for  {
			_, err := os.Stdin.Read(bb)
			if err != nil {
				panic("Could not read input.")
			}
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
				//cmd1.Process.Kill()
				//cmd2.Process.Kill()
				os.Exit(4)
			case "r":
				b.restart()
			}
		}
	}
}

func (b *board) logic() {
	if b.dir != stop {
		b.setSnake(b.dir)
	}
}

func (b *board) displayStage()  {
	utils.ClearTerminal()
	b.logic()

	utils.ClearTerminal()
	for i, y := range b.stage{
		for j := range y {
			if b.stage[i][j] == 1 && j == len(y) - 1{
				fmt.Println("*")
			} else if b.stage[i][j] == 1 && (i == 0 || i == len(y) - 1 || j == 0) {
				fmt.Print("*")
			} else if b.stage[i][j] == 1 {
				fmt.Print("*")
			} else if b.stage[i][j] == 2 {
				fmt.Print("o")
			} else if b.stage[i][j] == 3 {
				fmt.Print("x")
			}  else if b.stage[i][j] == 4 {
				fmt.Print("@")
			} else {
				fmt.Print(" ")
			}
		}
	}

	fmt.Println()
	fmt.Println()
	if !b.gameOver {
		fmt.Printf("%s: %d, %s: %d", "Height", b.height, "Width", b.width)
		fmt.Printf("\t\t\t\t\t%s\n\n","INSTRUCTIONS")
		fmt.Println("Current round: ", b.currentRound)
		fmt.Printf("\t\t\t\t\t\t\t%s\n","w: Move Up")
		fmt.Println("Score: ", b.score)
		fmt.Printf("\t\t\t\t\t\t\t%s\n","s: Move Down")
		fmt.Println("Snake length: ", len(b.snake.Tail) + 1)
		fmt.Printf("\t\t\t\t\t\t\t%s\n","a: Move Left")
		fmt.Printf("%s: %+v\n", "Snake Head Coordinates", *b.snake.Head)
		fmt.Printf("\t\t\t\t\t\t\t%s\n","d: Move Right")
	} else {
		fmt.Println("GAME OVER.")
		fmt.Println("PRESS 'r' to restart")
		fmt.Println("PRESS 'x' to End")
	}

	time.Sleep(100* time.Millisecond)
}

func createBoard(size int) *board {
	b := &board{
		width: size * 3,
		height: size,
		dir: up,
		stop: false,
	}
	b.setStage()
	return b
}

func (b *board) restart() {
	b.setStage()
	b.score = 0
	b.currentRound = 0
	b.stop = false
	b.gameOver = false
}

func (b *board) setStage() {
	b.stage = make([][]int, b.height)
	for i := range b.stage {
		b.stage[i] = make([]int, b.width)
	}

	b.snake = &snake{
		length: 1,
		Head: &coord{
			x: b.width/2,
			y: b.height/2,
		},
	}

	b.food = createFood(b.width, b.height)
	//food := CreateFood(b.width, b.height)
	//b.food = food

	b.setBorder()
	xFood, yFood := b.food.pos.get()
	b.stage[yFood][xFood] = 3
}

func (b *board) setBorder() {
	for i, y := range b.stage{
		for j := range y {
			if i == 0 || i == len(b.stage) -1 {
				b.stage[i][j] = 1
				continue
			} else if j == 0 || j == len(y) -1 {
				b.stage[i][j] = 1
				continue
			}
			b.stage[i][j] = 0
		}
	}
}
