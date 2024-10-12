package hw02unpackstring

import (
	"errors"
	"unicode"
)

var (
	ErrInvalidString         = errors.New("invalid string")
	ErrInvalidEscapeSequence = errors.New("invalid escape sequence")
)

func Unpack(str string) (string, error) {
	disallowNumber := true
	result := []rune{}
	literalMode := false

	for _, grapheme := range str {
		if literalMode {
			if unicode.IsLetter(grapheme) {
				return "", ErrInvalidEscapeSequence
			}
			result = append(result, grapheme)
			literalMode = false
			disallowNumber = false
			continue
		}

		if grapheme == '\\' {
			literalMode = true
			continue
		}

		if unicode.IsNumber(grapheme) {
			if disallowNumber {
				return "", ErrInvalidString
			}
			result = multiplyTailBy(result, int(grapheme-'0'))
			disallowNumber = true
			continue
		}

		if unicode.IsLetter(grapheme) {
			result = append(result, grapheme)
			disallowNumber = false
			continue
		}
		return "", ErrInvalidString
	}

	if literalMode {
		return "", ErrInvalidEscapeSequence
	}
	return string(result), nil
}

func multiplyTailBy(result []rune, n int) []rune {
	var toAppend rune
	result, toAppend = result[:len(result)-1], result[len(result)-1]
	for i := 0; i < n; i++ {
		result = append(result, toAppend)
	}
	return result
}
