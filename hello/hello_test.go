package main

import "testing"




func TestHello(t *testing.T) { 
	
	assertCorrectMessage := func (t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %s", got , want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) { 
		got := Hello("Chriss")
		want := "Hello, Chriss"

		assertCorrectMessage(t,got,want)
	})

	t.Run("Say hello world when the param is empty", func(t *testing.T) {
		got := Hello("")
		want := Hello("world")

		assertCorrectMessage(t,got,want)
	}) 
}

