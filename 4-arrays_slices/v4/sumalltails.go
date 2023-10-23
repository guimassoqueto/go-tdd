package main

func SumAllTails(numbers... []int) []int{
	var sum []int
	for _, numbers := range numbers {
		tail := numbers[1:]
		sum = append(sum, Sum(tail))
	}
	return sum
}