package main

import "testing"

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}
		got := Reduce[int]([]int{1,2,3}, multiply, 1)
		want := 6
		if  got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})


	t.Run("concatenate strings", func(t *testing.T) {
		concat := func(x, y string) string {
			return x + y
		}
		got := Reduce[string]([]string{"hello"," ","world"}, concat, "")
		want := "hello world"
		if  got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}