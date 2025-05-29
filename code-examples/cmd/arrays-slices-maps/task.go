package main

import (
	"fmt"
	"slices"
	"sort"
)

// When working with the code above, you have seen that there is a problem. Changes to section cause unwanted changes to inventory.
// You have been told that the slices package may contain a function to solve this issue so that changes can freely be made to section without affecting inventory

func main() {
	//slicesFunc()

	message := "the cat in the hat comes back"
	lCount := make(map[rune]uint8)

	// counts letters
	for _, ch := range message {
		if ch >= 'a' && ch <= 'z' {
			lCount[ch]++
		}
	}
	fmt.Println(lCount)

	// extracts letters (keys) into slice
	letters := make([]rune, 0, len(lCount))
	for ch := range lCount {
		letters = append(letters, ch)
	}
	fmt.Println(letters)

	// alphabetically sorts
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})
	fmt.Println(letters)

	for _, ch := range letters {
		fmt.Printf("%c: %d\n", ch, lCount[ch])
	}
}

func slicesFunc() {
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
