package main 

import "testing"

func TestRomanNumerals(t *testing.T) {
	t.Run("basic test", func(t *testing.T) {
		got := ConvertToRoman(1)
		want := "I"
	
		assertEqual(t, got, want)
	})
}

func assertEqual(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}