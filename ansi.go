package main

import (
	"fmt"
	"io"
	"os"
)

const (
	CSI string = "\x1b"
)

var (
	Black   = "\033[1;30m%s\033[0m"
	Red     = "\033[1;31m%s\033[0m"
	Green   = "\033[1;32m%s\033[0m"
	Yellow  = "\033[1;33m%s\033[0m"
	Purple  = "\033[1;34m%s\033[0m"
	Magenta = "\033[1;35m%s\033[0m"
	Teal    = "\033[1;36m%s\033[0m"
	White   = "\033[1;37m%s\033[0m"
)

func GetCursorPosition() int {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fmt.Printf("\x1b[6n")

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = stdout

	return int(out[2])
}

func ClearLine() {
	fmt.Printf("\x1b[2K")
}

func CursorBack() {
	fmt.Print("\x1b[0G")
}

func CursorBackN(n int) {
	fmt.Printf("\x1b[%dG", n)
}

func CursorBackOne() {
	fmt.Printf("\x1b[1D")
}

func CursorFwdOne() {
	fmt.Printf("\x1b[1C")
}

func ClearToEndOfLine() {
	fmt.Print("\x1b[K")
}

func PrintRed(s string) {
	fmt.Printf(Red, s)
}

func PrintGreen(s string) {
	fmt.Printf(Green, s)
}
