package main

import (
	"fmt"
	"math"
)

func main() {
	// can include statements before conditions in if statements i.e. if sval, err:= funcCall(); cerr != nil {...}
	var num float64
	fmt.Print("Enter number: ")
	fmt.Scanln(&num)

	number(num)
}

func number(num float64) {
	if num > 0 {
		if num == math.Trunc(num) && int(num)%2 == 0 { //math.Trunc checks if num is a whole number
			fmt.Println("Positive, Even")
		} else {
			fmt.Println("Positive, Odd")
		}
	} else if num < 0 {
		if num == math.Trunc(num) && int(num)%2 == 0 {
			fmt.Println("Negative, Even")
		} else {
			fmt.Println("Negative, Odd")
		}
	} else {
		fmt.Println("Zero, Even")
	}
}

// m := time.Now().Month()
// switch m {
// case time.September, time.April, time.June, time.November:
//   fmt.Println("There are 30 days in", m)
// case time.February:
//   fmt.Println("There are either 28 or 29 days in", m)
// default:
//   fmt.Println("There are 31 days in", m)
// }
