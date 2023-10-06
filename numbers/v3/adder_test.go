package main

import "testing"

func TestAdder(t *testing.T) {
	t.Run("calc the sum of two integers", func(t *testing.T) {
		got := Adder(2, 2)
		expect := 4

		assertCorrectValue(t, got, expect)
	})

	t.Run("calc the sum of three or more integers", func(t *testing.T) {
		got := Adder(3, 2, 6)
		expect := 11

		assertCorrectValue(t, got, expect)
	})

	t.Run("calc the sum of two or more floats", func(t *testing.T) {
		got := Adder(1.2, 5.3)
		expect := 6.5
		assertCorrectValue(t, got, expect)
	})

	t.Run("calc the sum of two or more floats", func(t *testing.T) {
		got := Adder(1, 5.3)
		expect := 6.3
		assertCorrectValue(t, got, expect)
	})
}

func assertCorrectValue[T int|float64](t testing.TB, got, expect T) {
	t.Helper()
	if got != expect {
		t.Errorf("got %v want %v", got, expect)
	}
}