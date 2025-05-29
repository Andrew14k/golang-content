package main

import "fmt"

// functions can take params, take no params and return, take params and return, return multiple things

////////////////////////////////////////////////////////

type average func(...int) float64 //declaring average as a func type

func meanAverage(nums ...int) float64 { //so takes parameter a slice of whatever size of ints
	var avg float64
	for _, v := range nums {
		avg += float64(v)
	}
	return avg / float64(len(nums))
}

func display(op average, nums ...int) {
	fmt.Println(op(nums...))
}

////////////////////////////////////////////////////////

type operation func(a, b int) int

func calculate(op operation, a, b int) int {
	return op(a, b)
}

////////////////////////////////////////////////////////

func makeCounter() func() int { //state maintained within function
	count := 0

	return func() int {
		count++
		return count
	}
}

////////////////////////////////////////////////////////

func main() {
	// fmt.Println(meanAverage(1, 2, 3, 4, 5, 6, 6, 5, 4, 3, 3, 7))

	// nums := []int{1, 2, 3, 4, 5, 5, 3, 2}
	// fmt.Println(meanAverage(nums...))

	// m := meanAverage
	// fmt.Println(m(2, 3))

	// //var safeAvg average = meanAverage //another way to assign function - good for readability and type safety
	// display(m, nums...)

	//anonymous function passed to calculate
	result := calculate(func(a, b int) int {
		return a * b
	}, 4, 5)
	fmt.Println(result)

	//closures - extrememl useful
	countA := makeCounter()
	countB := makeCounter()
	fmt.Printf("Counter A: %d\n", countA())
	fmt.Printf("Counter A: %d\n", countA())
	fmt.Printf("Counter B: %d\n", countB())
	fmt.Printf("Counter A: %d\n", countA())
	fmt.Printf("Counter B: %d\n", countB())
	fmt.Printf("Counter A: %d\n", countA())
	fmt.Printf("Counter B: %d\n", countB())
}

//first class citizens means you can treat a function as a type
