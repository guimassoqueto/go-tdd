package main 

import "testing"

func TestSum(t *testing.T) {
	t.Run("calculate the sum of an slice of any size", func(t *testing.T) {
		numbers := []int{1,2,3,4,5,6,7}
		got := Sum(numbers)
		expect := 28

		if got != expect {
			t.Errorf("expect '%d' got '%d'", expect, got)
		}
	})
}