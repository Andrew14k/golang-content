package main

import "fmt"

// issues with having to repeat similar codes for different data type cases (for examples)
// generics allow functions to receive additional paramters

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumSlice[float32](float32Slice))
}

// T = type (so could have been int before, for example)
func sumSlice[T int | float32 | float64](slice []T) T { //passing types as additional params in [] brackets
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// any type = [T any] can be used for cases of if slices are empty (for example)
// any generic cannot be used in the case above for example, as sum += v is not going to work for every type, such as bools
