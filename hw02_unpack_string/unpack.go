package hw02unpackstring

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
	disallowNumber := true
	result := []string{}
	literalMode := false

	graphemes := uniseg.NewGraphemes(str)

	for graphemes.Next() {
		grapheme := graphemes.Str()
		graphemeType := parseGrapheme(grapheme)

		if literalMode {
			if graphemeType.isGrapheme {
				return "", ErrInvalidEscapeSequence
			}
			result = append(result, grapheme)
			literalMode = false
			disallowNumber = false
			continue
		}

		if graphemeType.isEscape {
			literalMode = true
			continue
		}

		if graphemeType.isNumber {
			if disallowNumber {
				return "", ErrInvalidString
			}
			result = multiplyTailBy(result, graphemeType.number)
			disallowNumber = true
			continue
		}

		result = append(result, grapheme)
		disallowNumber = false
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

type graphemeType struct {
	isNumber   bool
	number     int
	isGrapheme bool
	isEscape   bool
}

func parseGrapheme(grapheme string) graphemeType {
	if grapheme == `\` {
		return graphemeType{isEscape: true}
	}
	n, err := strconv.Atoi(grapheme)
	if err == nil {
		return graphemeType{isNumber: true, number: n}
	}
	return graphemeType{isGrapheme: true}
}
