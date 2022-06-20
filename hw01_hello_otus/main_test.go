package main

import "testing"

func TestReverseString(t *testing.T) {
	got := ReverseString("Hello, OTUS!")

	if got != "!SUTO ,olleH" {
		t.Errorf("got %q, wanted %q", got, "!SUTO ,olleH")
	}
}
