package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	runeSlice := []rune(str)
	ln := len(runeSlice)
	if ln == 0 {
		return "", nil
	}
	if unicode.IsDigit(runeSlice[0]) {
		return "", ErrInvalidString
	}
	var sb strings.Builder
	for i := 0; i < ln; i++ {
		if i+1 < ln {
			if unicode.IsDigit(runeSlice[i]) && unicode.IsDigit(runeSlice[i+1]) {
				return "", ErrInvalidString
			}
			// Checking zero exist in the current and next rune.
			if runeSlice[i+1] == 48 || runeSlice[i] == 48 {
				continue
			} else if unicode.IsDigit(runeSlice[i]) {
				sb.WriteString(strings.Repeat(string(runeSlice[i-1]), int(runeSlice[i]-'0')-1))
				continue
			}
		}
		if runeSlice[i] != 48 {
			sb.WriteString(string(runeSlice[i]))
		}
	}
	return sb.String(), nil
}
