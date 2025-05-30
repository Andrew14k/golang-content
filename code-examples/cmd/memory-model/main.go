package main

import "fmt"

// go uses stacks and heaps
//ususally stuff goes to stack - very quick
func main() {
	var x int  //int value
	var y *int //pointer to an int value

	//zero value of y as it is a pointer is nil

	x = 42
	y = &x //get pointer to x
	//*y = 8 - changes x by derefernecing y

	fmt.Printf("%d: %p", *y, y) // returns 42: 0xc00000a0d8	//y* is used here to dereference, so returns the actual value being pointed to (42)
}

// func foo() *int {
// 	x := 42
// 	y := &x
// 	return y //returns 0xc00000a0d8 - memory location of x due to *int return in function statement
// }

///////////////////////////////////////////////////////////////////////////////////////////////////////
/*
Stack: - fast local storage for immediate use of memory - data not being returned
	x = 42
	y = pointer of x (where x exists in memory) - so is pushed to the heap
	if function wasnt return y, it would go on stack. becuase its being returned, it is 'escaping' the function, so stores it on the heap

Heap: - managed piece of memoery for longer term - data being returned
	42 (int) - stored longer term as it cannot be forgotten after function is executed (return statement)
*/

/*
Pointer types denoted by *
*/
