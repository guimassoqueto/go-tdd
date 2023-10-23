package main 

import "testing"

func TestSum(t *testing.T) {
	t.Run("calculate the sum of every element in the slice", func(t *testing.T) {
		numbers := [5]int{1,2,3,4,5}
		got := Sum(numbers)
		expect := 15

		if got != expect {
			t.Errorf("expect '%d' got '%d'", expect, got)
		}
	})
}
