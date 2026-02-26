package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func (t *testing.T){
		got := Hello("Anna", "")
		want := "Hello, Anna"
		assertMessage(t, got, want)	

	})
	t.Run("Empty input returns hello world", func(t *testing.T){
		got  := Hello("", "")
		want := "Hello, World"
	    assertMessage(t, got, want)	
	})
	t.Run("Use Spanish greeting", func(t *testing.T){
		got := Hello("Imanol", "spanish")
		want := "Hola, Imanol"
	    assertMessage(t, got, want)	
	})
	t.Run("Use French greeting", func(t *testing.T){
		got := Hello("Adrien", "french")
		want := "Bonjour, Adrien"
	    assertMessage(t, got, want)	
	})
	t.Run("Use English greeting", func(t *testing.T){
		got := Hello("Imanol", "english")
		want := "Hello, Imanol"
	    assertMessage(t, got, want)	
	})
	t.Run("Use weird case language", func(t *testing.T){
		got := Hello("Imanol", "French")
		want := "Bonjour, Imanol"
	    assertMessage(t, got, want)	
	})
}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}

}
