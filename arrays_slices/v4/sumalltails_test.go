package main

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	t.Run("sum the tails of arrays", func(t *testing.T) {
		got := SumAllTails([]int{1,2}, []int{0,9}, []int{3,4,6})
		want := []int{2,9,10}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	})
}