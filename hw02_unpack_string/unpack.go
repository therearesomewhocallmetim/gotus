package hw02unpackstring

// package main

import (
	"errors"
	"strconv"
	"strings"

	"github.com/rivo/uniseg"
)

var ErrInvalidString = errors.New("invalid string")
var ErrInvalidEscapeSequence = errors.New("invalid escape sequence")

func Unpack(str string) (string, error) {
	canAcceptNumber := false
	result := []string{}
	literalMode := false

	graphemes := uniseg.NewGraphemes(str)

	for graphemes.Next() {
		grapheme := graphemes.Str()
		if grapheme == "\\" {
			if literalMode {
				result = append(result, grapheme)
				literalMode = false
				continue
			}
			literalMode = true
			continue
		}
		n, err := strconv.Atoi(grapheme)
		if err != nil { // the grapheme is not a number
			if literalMode {
				return "", ErrInvalidEscapeSequence
			}
			result = append(result, grapheme)
			canAcceptNumber = true
			continue
		}
		if !canAcceptNumber {
			return "", ErrInvalidString
		}
		if literalMode {
			result = append(result, grapheme)
			literalMode = false
			continue
		}
		result = multiplyTailBy(result, n)
		canAcceptNumber = false

	}
	if literalMode {
		return "", ErrInvalidEscapeSequence
	}
	return strings.Join(result, ""), nil
}

func multiplyTailBy(result []string, n int) []string {
	var toAppend string
	result, toAppend = result[:len(result)-1], result[len(result)-1]
	for i := 0; i < n; i++ {
		result = append(result, toAppend)
	}
	return result
}
