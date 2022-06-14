package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	// Place your code here.
	runeSlice := []rune(str)
	var sb strings.Builder
	ln := len(runeSlice)
	if ln == 0 {
		return "", nil
	}
	if unicode.IsDigit(runeSlice[0]) {
		return "", ErrInvalidString
	}
	for i := 0; i < ln; i++ {
		if i+1 < ln {
			if unicode.IsDigit(runeSlice[i]) && unicode.IsDigit(runeSlice[i+1]) {
				return "", ErrInvalidString
			}
			if runeSlice[i+1] == 48 || runeSlice[i] == 48 {
				continue
			} else if unicode.IsDigit(runeSlice[i]) {
				sb.WriteString(strings.Repeat(string(runeSlice[i-1]), int(runeSlice[i]-'0')-1))
				continue
			}
		}
		sb.WriteString(string(runeSlice[i]))
	}

	return sb.String(), nil
}

func main() {
	line, _ := Unpack("d\n5abc")
	fmt.Print(line)
}
