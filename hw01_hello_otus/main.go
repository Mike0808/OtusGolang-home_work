package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	// Place your code here.
	fmt.Print(ReverseString("Hello, OTUS!"))
}

func ReverseString(s string) string {
	return stringutil.Reverse(s)
}
