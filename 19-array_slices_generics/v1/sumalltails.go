package main

func SumAllTails(numbers... []int) []int{
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}
	return Reduce[[]int](numbers, sumTail, []int{})
}