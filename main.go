package main

import (
	"fmt"
	"rollez/keys"
	"runtime"
)

func main() {

	os := runtime.GOOS

	fmt.Println("OS = ", os)

	terminal := keys.OpenTerminal()

	defer terminal.Close()

	fmt.Println("I will repeate whatever you say until you hit [ESC]")
	fmt.Println()

	for {

		line := terminal.ReadLine()

		PrintGreen("You Said: ")
		PrintRed(line)
		fmt.Println()
	}

}
