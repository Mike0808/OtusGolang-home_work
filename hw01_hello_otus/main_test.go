package main

import "testing"

func TestReverseString(t *testing.T) {

	got := ReverseString("Hello, OTUS!")
	want := "!SUTO ,olleH"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
