package main

import (
	"fmt"
	"slices"
)

// When working with the code above, you have seen that there is a problem. Changes to section cause unwanted changes to inventory.
// You have been told that the slices package may contain a function to solve this issue so that changes can freely be made to section without affecting inventory

func main() {
	inventory := []int{100, 200, 300, 400, 500, 600}
	fmt.Println("Original inventory:", inventory)

	section := slices.Clone(inventory[2:5]) //clone function makes a copy of a slice
	fmt.Println("Section before modification:", section)

	// section := make([]int, 3)
	// copy(section, inventory[2:5])
	section[0] = 999

	fmt.Println("Modified section:", section)
	fmt.Println("Inventory after section change:", inventory)
}
