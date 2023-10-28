package main

func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce[int](numbers, add, 0)
}