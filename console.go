package main

import (
	"fmt"
	"os"
	"rollez/keys"

	"github.com/eiannone/keyboard"
)

type Console struct {
	cursor   int
	line     []rune
	history  History
	terminal keys.Terminal
}

func MakeConsole() Console {
	pTerminal := keys.OpenTerminal()
	return Console{cursor: 0, line: make([]rune, 255), terminal: *pTerminal}
}

func (c *Console) CursorRight() {
	c.cursor += 1
	CursorFwdOne()
}

func (c *Console) CursorLeft() {
	// fmt.Printf("cursor=%d  prompt=%d", c.cursor, len(c.prompt))

	c.cursor -= 1
	CursorBackOne()

}

func (c *Console) HandleBackspace() {

	c.cursor -= 1
	c.removeLastChar()
	CursorBackOne()
	ClearToEndOfLine()

}

func (c *Console) PrintChar(char rune) {
	c.line[c.cursor] = char
	c.cursor += 1
	fmt.Print(string(char))
}

func (c *Console) HandleEnter() {
	c.history.Push(string(c.line))
}

func (c *Console) GetLine() string {

	stillTyping := true

	for stillTyping {
		stillTyping = c.HandleKeyStroke()
	}

	c.history.ResetIndex()
	lineToReturn := c.line
	c.line = make([]rune, 255)
	return string(lineToReturn)

}

func (c *Console) ResetLine() {
	ClearLine()
	fmt.Print(string(c.line))
	CursorBackN(1)
}

func (c *Console) HandleArrowUp() {
	prev, ok := c.history.Pop()
	if ok {
		c.line = []rune(prev)
		c.ResetLine()
	}

}

func (c *Console) HandleDownArrow() {
	next, ok := c.history.Next()
	if ok {
		c.line = []rune(next)
	} else {
		c.line = make([]rune, 0)
	}
	c.ResetLine()
}

func (c *Console) HandleKeyStroke() bool {

	char, key, err := keyboard.GetSingleKey()

	if err != nil {
		panic(err)
	}

	switch key {
	case keyboard.KeyEsc:
		os.Exit(0)
	case keyboard.KeyArrowUp:
		c.HandleArrowUp()
	case keyboard.KeyArrowDown:
		c.HandleDownArrow()
	case keyboard.KeyArrowRight:
		c.CursorRight()
	case keyboard.KeyArrowLeft:
		c.CursorLeft()
	case keyboard.KeyBackspace:
		c.HandleBackspace()
	case keyboard.KeyBackspace2:
		c.HandleBackspace()
	case keyboard.KeySpace:
		c.PrintChar(' ')
	case keyboard.KeyEnter:
		c.HandleEnter()
		return false
	default:
		c.PrintChar(char)
	}

	return true
}

func (c *Console) removeLastChar() {

	if len(c.line) > 0 {
		c.line = c.line[:len(c.line)-1]

	}
}
