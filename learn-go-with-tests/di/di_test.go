package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "Emen")

	got := buffer.String()
	want := "Hello, Emen"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
