package utils

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

var clear map[string]func()

func RegisterUtils()  {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ClearTerminal()  {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Unsupported platform.")
	}
}

func ConvertToUnicode(s string) string {
	r, err := strconv.ParseInt(strings.TrimPrefix(s, "\\U"), 16, 32)
	if err != nil {
		return "*"
	}
	return string(r)
}