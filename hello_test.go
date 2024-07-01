package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Frankie")
		want := "Hello, Frankie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
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
