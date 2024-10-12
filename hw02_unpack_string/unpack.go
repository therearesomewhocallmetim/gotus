package hw02unpackstring

// package main

import (
	"errors"
	"strconv"
	"strings"

	"github.com/rivo/uniseg" //nolint:depguard
)

var (
	ErrInvalidString         = errors.New("invalid string")
	ErrInvalidEscapeSequence = errors.New("invalid escape sequence")
)

func Unpack(str string) (string, error) {
	canAcceptNumber := false
	result := []string{}
	literalMode := false

	graphemes := uniseg.NewGraphemes(str)

	for graphemes.Next() {
		grapheme := graphemes.Str()
		graphemeType := parseGrapheme(grapheme)

		if literalMode {
			if graphemeType.IsGrapheme {
				return "", ErrInvalidEscapeSequence
			}
			result = append(result, grapheme)
			literalMode = false
			canAcceptNumber = true
			continue
		}

		if graphemeType.IsEscape {
			literalMode = true
			continue
		}

		if graphemeType.IsNumber {
			if !canAcceptNumber {
				return "", ErrInvalidString
			}
			result = multiplyTailBy(result, graphemeType.Number)
			canAcceptNumber = false
			continue
		}

		result = append(result, grapheme)
		canAcceptNumber = true
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

const Escape = `\`

type GraphemeType struct {
	IsNumber   bool
	Number     int
	IsGrapheme bool
	// Grapheme   string
	IsEscape bool
}

func parseGrapheme(grapheme string) GraphemeType {
	if grapheme == `\` {
		return GraphemeType{IsEscape: true}
	}
	n, err := strconv.Atoi(grapheme)
	if err == nil {
		return GraphemeType{IsNumber: true, Number: n}
	}
	return GraphemeType{IsGrapheme: true}
}
