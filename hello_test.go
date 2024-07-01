package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people in english", func(t *testing.T) {
		got := Hello("Frankie", "English")
		want := "Hello, Frankie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Mobey", "French")
		want := "Bonjour, Mobey"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when empty args are supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// testing.TB allows calling helper function from a test, or a benchmark
	t.Helper()
	// Helper indicates this is not a test or code being tested
	// the line number reported will be in the function call rather than inside our test helper

	if got != want {
		t.Errorf("got %q want %q", got, want)
		t.Fail()
	}
}
