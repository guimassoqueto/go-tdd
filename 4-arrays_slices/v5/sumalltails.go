package main

func SumAllTails(numbers... []int) []int{
	var sum []int
	for _, numbers := range numbers {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			tail := numbers[1:]
			sum = append(sum, Sum(tail))
		}
	}
	return sum
}