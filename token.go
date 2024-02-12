package main

import (
	"fmt"
	"regexp"
)

type TokenType int

const (
	Number TokenType = iota
	Dice
	Operation
	Error
)

type Token struct {
	raw       string
	tokenType TokenType
}

func validateTokenSlice(tokens []Token) bool {

	for i, t := range tokens {

		if t.tokenType == Error { // if any Error types fail early
			fmt.Printf("Error = %q \n", t.raw)
			return false
		}

		if i%2 == 0 {
			if t.isOperationToken() { // cannot start with op
				return false
			}
		} else {
			if t.isValueToken() { // two values in a row is invalid
				return false
			}
		}

	}

	if tokens[len(tokens)-1].isOperationToken() { // no trailing ops
		return false
	}

	return true
}

func GetTokenType(token string) TokenType {

	if isNumber(token) {
		return Number
	}

	if isOp(token) {
		return Operation
	}

	if isDice(token) {
		return Dice
	}

	return Error
}

func (t *Token) isValueToken() bool {
	return t.tokenType == Dice || t.tokenType == Number
}

func (t *Token) isOperationToken() bool {
	return t.tokenType == Operation
}

func isNumber(token string) bool {
	match, _ := regexp.MatchString(`^\d+$`, token)
	return match
}

func isOp(token string) bool {
	match, _ := regexp.MatchString(`[-+%*]`, token)
	return match
}

func isDice(token string) bool {
	match, _ := regexp.MatchString(`[1-9]\d*d[1-9]\d*`, token)
	return match
}

func (tt TokenType) String() string {
	return [...]string{"Number", "Dice", "Operation", "Error"}[tt]
}
