package main

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{height: 10.0, width: 10.0}
		got := rectangle.Perimeter()
		want := 40.0
		assertEqual(t, got, want)
	})
}

func assertEqual(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}