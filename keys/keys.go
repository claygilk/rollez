package keys

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"

	"golang.org/x/sys/unix"
)

var (
	ESC rune = 27
)

type Terminal struct {
	input   int
	output  *os.File
	reader  *bufio.Reader
	termios unix.Termios
}

func (t *Terminal) Close() {
	t.termios.Lflag |= unix.ECHO                           // set 8th bit to 1
	unix.IoctlSetTermios(t.input, unix.TCSETS, &t.termios) // make syscall
}

func OpenTerminal() *Terminal {

	input, _ := unix.Open("/dev/tty", unix.O_RDONLY, 0)

	termios, _ := unix.IoctlGetTermios(input, unix.TCGETS)
	termios.Lflag &^= unix.ECHO   // Don't print character
	termios.Lflag &^= unix.ICANON // Don't buffer until new line
	unix.IoctlSetTermios(input, unix.TCSETS, termios)

	terminal := Terminal{input: input, termios: *termios}

	return &terminal
}

func (t *Terminal) ReadLine() string {

	lineOfRunes := make([]rune, 100)

	for {

		char, _ := t.ReadOneRune()

		if char != '\n' {
			fmt.Print(string(char))
			lineOfRunes = append(lineOfRunes, char)
		} else {
			fmt.Println()
			break
		}

	}

	return string(lineOfRunes)

}

func (t *Terminal) ReadOneRune() (rune, int) {

	for {

		buf := make([]byte, 128)
		bytesRead, _ := unix.Read(t.input, buf)

		if bytesRead > 0 {

			data := make([]byte, bytesRead)
			copy(data, buf)

			r, rBytes := utf8.DecodeRune(data)

			if r == ESC {
				t.Close()
				os.Exit(0)
			}

			return r, rBytes
		}
	}
}
