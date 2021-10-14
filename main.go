package main

import (
	"github.com/Kolakanmi/snake/game"
	"github.com/Kolakanmi/snake/utils"
)

func main()  {
	utils.RegisterUtils()
	board := game.CreateBoard(20)
	board.Run()
}
