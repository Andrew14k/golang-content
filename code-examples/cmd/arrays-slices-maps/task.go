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

	messageMap()
}

func messageMap() {
	message := "the cat in the hat comes back"
	lCount := make(map[rune]uint8)

	// counts letters
	for _, ch := range message {
		if ch >= 'a' && ch <= 'z' { //ignores whitespace
			lCount[ch]++
		}
	}
	//fmt.Println(lCount) //map[97:3 98:1 99:3 101:3 104:3 105:1 107:1 109:1 110:1 111:1 115:1 116:4] unicode code points

	// creates a slice that takes the keys of lCount
	letterKey := []rune{}
	for ch := range lCount {
		letterKey = append(letterKey, ch)
	}
	//fmt.Println(letterKey) //[116 104 101 99 97 105 110 115 111 109 98 107] or other order

	// alphabetically sorts - func(i, j) how elements are compared
	sort.Slice(letterKey, func(i, j int) bool { return letterKey[i] < letterKey[j] })
	//fmt.Println(letterKey) //[97 98 99 101 104 105 107 109 110 111 115 116]

	for _, ch := range letterKey {
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
