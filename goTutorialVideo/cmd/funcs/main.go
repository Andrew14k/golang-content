package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe()

	var printValue string = "World"
	printStr(printValue)

	var numerator int = 11
	var denominator int = 0
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("Result of integer division is %v with remainder %v", result, remainder)
	}

	//with switch cases, write switch{ case ...: case ... : default:... } break is implied, doesnt need to be written
	switch remainder {
	case 0:
		fmt.Printf("dividion was exact")
	case 1, 2:
		fmt.Printf("division was close")
	default:
		fmt.Printf("divison was not close")
	}
}

func printMe() {
	fmt.Println("Hello")
}

func printStr(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
