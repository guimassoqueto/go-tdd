package main

import "testing"

func TestArea(t *testing.T) {
	t.Run("test the right calculation of the perimeter", func(t *testing.T) {
		got := Area(10.0, 10.0)
		want := 100.0
		assertEqual(t, got, want)
	})
}