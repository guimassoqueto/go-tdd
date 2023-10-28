package main

func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
	result := initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}