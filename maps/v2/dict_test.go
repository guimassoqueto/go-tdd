package main

import "testing"

func TestSearch(t *testing.T) {
	dict := Dict{"test": "this is just a test"}	

	t.Run("should return the expected key", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_ , err := dict.Search("not-valid")
		want := "could not find the word you were looking for"
		
		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}