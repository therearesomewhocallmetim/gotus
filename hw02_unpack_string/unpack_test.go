package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require" //nolint:depguard
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		// uncomment if task with asterisk completed
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		{input: `qwe3\3`, expected: `qweee3`},
		{input: `qw,3\3`, expected: `qw,,,3`},
		{input: `qw 3\3`, expected: `qw   3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{
		"3abc",
		"45",
		"aaa10b",
		"d̷̙͈̙̓͛̇ͨ̾̐4   i̧̻̪͍͒ͬ͆̀̒͝      a̷̯̞̻ͤ͏͙͙̜͘     c̨̭͚̼̙̍ͨ͌̚͝      r̨͇̯̱̍͋ͧͩ̕͜    i̭͍̘̞̣̱̐͛͜͟     t̢͇̙̯ͭ̐̇̽ͥ̈    三ï̴̮̺̜̙̪͉̩̮    c̰̞͍͖̪̣ͮ̂̚͝     ì̧̙̤́͂̽̑̚͡    s̗͕̘̩̯̎́̚̕͟      ṁ̫̩̭̊͐ͪ́̈͢   ", //nolint:lll
	}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestUnpackInvalidEscapeSequence(t *testing.T) {
	invalidStrings := []string{`\`, `hello\d`, `ab\,`}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidEscapeSequence), "actual error %q", err)
		})
	}
}
