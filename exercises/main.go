package main

import (
	"fmt"
	"strconv"
)

func main() {
	fizzBuzz()
	fmt.Println(" ")

	var input string
	fmt.Print("input for type assessment: ")
	fmt.Scanln(&input)
	typeConversion(input)
	fmt.Println(" ")

	prime := isPrime(5)
	fmt.Println(prime)
	fmt.Println(" ")

	numbers := []int{2, 4, 6, 3, 9, 5}
	fmt.Println(arrayAnalysis(numbers))
	fmt.Println(" ")

	word := "hello"
	fmt.Println(reverseString(word))
}

// fizz buzz iterator
func fizzBuzz() {
	count := 0
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
			count++
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		}
	}
	fmt.Println(count)
}

// check type for input
func typeConversion(input string) {
	if i, err := strconv.Atoi(input); err == nil {
		fmt.Println("input is integer: ", i)
		return
	}

	if f, err := strconv.ParseFloat(input, 64); err == nil {
		fmt.Println("input is float: ", f)
		return
	}

	if b, err := strconv.ParseBool(input); err == nil {
		fmt.Println("input is a boolean: ", b)
		return
	}

	if len(input) == 0 {
		fmt.Println("input is empty or unrecognized")
		return
	}

	fmt.Println("input is a string:", input)
}

// determine if number is prime
func isPrime(num int) bool {
	var result bool = true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			result = false
		}
	}
	return result
}

// return min, max and average of array
func arrayAnalysis(numbers []int) (min int, max int, average float64) {
	min, max = numbers[0], numbers[0]
	var sum int

	if len(numbers) == 0 {
		return 0, 0, 0
	}

	for _, i := range numbers {
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
		sum += i
	}

	average = float64(sum) / float64(len(numbers))
	return min, max, average
}

func reverseString(s string) string {
	//rune represents a unicode paint, i.e character
	runes := []rune(s)                                    //converts string into slice of characters/runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 { //swaps elements
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
