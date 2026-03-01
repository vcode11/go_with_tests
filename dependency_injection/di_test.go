package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "alice")
	got := buffer.String()
	want := "Hello, alice"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
