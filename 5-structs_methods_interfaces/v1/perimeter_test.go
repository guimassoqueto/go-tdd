package main

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("test the right calculation of the perimeter", func(t *testing.T) {
		got := Perimeter(10.0,10.0)
		want := 40.0
		assertEqual(t, got, want)
	})
}

func assertEqual(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}