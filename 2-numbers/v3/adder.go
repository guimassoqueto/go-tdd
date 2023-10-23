package main

func Adder[T int|float64](nums... T) T {
	var total T
	for _, v := range nums {
		total += v
	}
	return total
}

func main() {
	print(Adder(2,3))
}
