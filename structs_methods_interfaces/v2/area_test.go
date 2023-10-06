package main

import "testing"

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{width: 12, height: 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{radius: 10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}