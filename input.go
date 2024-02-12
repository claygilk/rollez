package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func ScanLine() string {

	scanner.Scan()

	err := scanner.Err()

	if err != nil {
		fmt.Println(err.Error())
	}

	return scanner.Text()
}
