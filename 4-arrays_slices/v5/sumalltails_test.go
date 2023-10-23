package main

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	t.Run("sum the tails of arrays", func(t *testing.T) {
		got := SumAllTails([]int{1,2}, []int{0,9}, []int{3,4,6})
		want := []int{2,9,10}

		assertDeepEqual(t, got, want)
	})

	t.Run("secure calc the sum if the slice provided is empty", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0,9}, []int{3,4,6})
		want := []int{0,9,10}

		assertDeepEqual(t, got, want)
	})
}

func assertDeepEqual(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}