package main

import (
	"fmt"
	"rollez/keys"
)

func main() {

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
