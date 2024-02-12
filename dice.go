package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

type Roll struct {
	numberOfDice int
	sizeOfDice   int
	total        int
	diceRolls    []int
}

func RollDice(diceToken string) Roll {

	split := strings.Split(diceToken, "d")
	log.Default().Print(split)
	nDice, _ := strconv.Atoi(split[0])
	sDice, _ := strconv.Atoi(split[1])

	r := Roll{numberOfDice: nDice, sizeOfDice: sDice, diceRolls: make([]int, nDice)}
	fmt.Println(r)
	r.doRoll()

	return r

}

func (r *Roll) doRoll() {
	for i := 0; i < r.numberOfDice; i++ {
		d := rand.Intn(r.sizeOfDice) + 1
		r.total += d
		r.diceRolls[i] = d
	}
}
