//GENERICS // - functions written to work on multiple types
// func Index[T comparable](s []T, x T) int
// above means that s is a slice of any type T that fulfills the built-in constraint comparable.
// x is a value of the same type

//comparable makes it possible to use the == and != operators on values of the type.

package main

import "fmt"

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}

// si returns 2, ss returns -1

// list represents a singly-linked list that holds values of any type.
// type List[T any] struct {
// 	next *List[T]
// 	val  T
// }
