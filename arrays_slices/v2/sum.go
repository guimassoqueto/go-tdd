package main

func Sum(numbers [5]int) int {
	var sum int
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func main() {}