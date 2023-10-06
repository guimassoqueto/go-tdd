package main

import (
	"math/rand"
	"testing"
)

func randInt() int {
	return 1 + rand.Intn(10)
}

func TestRepeat(t *testing.T) {
	t.Run("loop 3 times", func(t *testing.T) {
		got := Repeat("a", 3)
		expected := "aaa"
		if got != expected {
			t.Errorf("expect %q got %q", got, expected)
		}
	})

	t.Run("loop 0 times", func(t *testing.T) {
		got := Repeat("a", 0)
		expected := ""
		if got != expected {
			t.Errorf("expect %q got %q", got, expected)
		}
	})

	t.Run("loop -3 times", func(t *testing.T) {
		got := Repeat("a", -3)
		expected := ""
		if got != expected {
			t.Errorf("expect %q got %q", got, expected)
		}
	})
}

func BenchmarkRepeat(B *testing.B) {
	for i := 0; i < B.N; i++ {
		Repeat("a", randInt())
	}
}