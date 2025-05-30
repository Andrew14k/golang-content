package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// fizzBuzz()
	// fmt.Println(" ")

	// var input string
	// fmt.Print("input for type assessment: ")
	// fmt.Scanln(&input)
	// typeConversion(input)
	// fmt.Println(" ")

	// prime := isPrime(5)
	// fmt.Println(prime)
	// fmt.Println(" ")

	// numbers := []int{2, 4, 6, 3, 9, 5}
	// fmt.Println(arrayAnalysis(numbers))
	// fmt.Println(" ")

	// word := "hello"
	// fmt.Println(reverseString(word))

	//multiplicationTable(4)
	// arr := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(reverseArray(arr))

	//fmt.Println(removeIndex([]int{10, 20, 30, 40}, 2))

	fmt.Println(countWords([]string{"apple", "banana", "apple", "orange", "banana", "apple"}))
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

// prints multiplication table for each integer up to num, and show the multiples of each int up to num up to the numth integer
func multiplicationTable(num int) {
	for i := 1; i <= num; i++ {
		for x := 1; x <= num; x++ {
			fmt.Printf("%d ", i*x)
		}
		fmt.Println(" ") //splits for each num
	}
}

// reverse an array
func reverseArray(arr [5]int) [5]int {
	var reverse [5]int
	for i := 0; i < len(arr); i++ {
		reverse[i] = arr[len(arr)-1-i]
	}
	return reverse
}

// remove item at certain index of slice
func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

// count words frequency in string, return map
func countWords(text []string) map[string]int {
	sliceMap := make(map[string]int)

	singleString := strings.ToLower(strings.Join(text, " "))

	words := strings.Fields(singleString) //splits words by whitespace

	for _, word := range words {
		sliceMap[word]++
	}
	return sliceMap
}
