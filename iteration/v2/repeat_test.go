package main

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	expected := "aaaaa"
	if got != expected {
		t.Errorf("expect %q got %q", got, expected)
	}
}

func BenchmarkRepeat(B *testing.B) {
	for i := 0; i < B.N; i++ {
		Repeat("a")
	}
}