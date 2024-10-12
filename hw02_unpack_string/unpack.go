package hw02unpackstring

// package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/rivo/uniseg"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	// var lastChar string
	canAcceptNumber := false
	result := []string{}
	// candidate := ""

	graphemes := uniseg.NewGraphemes(str)

	for graphemes.Next() {
		grapheme := graphemes.Str()
		n, err := strconv.Atoi(grapheme)
		if err != nil { // the grapheme is not a number
			result = append(result, grapheme)
			canAcceptNumber = true
			continue
		}
		if !canAcceptNumber {
			fmt.Printf("We need char! %v, %d, %s", result, n, grapheme)
			return "", ErrInvalidString
		}
		var toAppend string
		result, toAppend = result[:len(result)-1], result[len(result)-1]
		for i := 0; i < n; i++ {
			result = append(result, toAppend)
		}
		canAcceptNumber = false

	}
	return strings.Join(result, ""), nil
}

// func main() {

// 	s, err := Unpack("d̷̙͈̙̓͛̇ͨ̾̐4   i̧̻̪͍͒ͬ͆̀̒͝      a̷̯̞̻ͤ͏͙͙̜͘     c̨̭͚̼̙̍ͨ͌̚͝      r̨͇̯̱̍͋ͧͩ̕͜    i̭͍̘̞̣̱̐͛͜͟     t̢͇̙̯ͭ̐̇̽ͥ̈    三ï̴̮̺̜̙̪͉̩̮    c̰̞͍͖̪̣ͮ̂̚͝     ì̧̙̤́͂̽̑̚͡    s̗͕̘̩̯̎́̚̕͟      ṁ̫̩̭̊͐ͪ́̈͢   ")
// 	if err != nil {
// 		panic("Hell")
// 	}
// 	println(s)
// }
