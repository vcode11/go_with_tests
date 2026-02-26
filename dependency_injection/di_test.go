package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T){
	buffer := bytes.Buffer{}
	Greet(&buffer, "Vishal")
	got := buffer.String()
	want := "Hello, Vishal"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}