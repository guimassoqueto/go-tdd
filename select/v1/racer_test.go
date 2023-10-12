package main

import (
	"testing"
)

func TestRacer(t *testing.T) {
	t.Run("", func(t *testing.T) {
		slowURL := "https://www.facebook.com.br"
		fastURL := "https://www.quii.dev"
		want := fastURL
		got := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}