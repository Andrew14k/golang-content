package main

import (
	"fmt"
	"math"
)

// fmt.Println("random number is ", rand.Intn(10))
// math.Sqrt(7)
// methods of packages must have a starting capital letter, meaning they are not exported/accessible if not capital
// type comes after variable name

// named return values are treated as variables and defined at the top of the function
// return statement without args returns the named return vales - known as a naked return
// only use naked returns in short functions

// var statements can be at package or function level
// := is used in place of var declaration with implicit type - only available within functions
// rune - unicode code point

// var i int
// j := i // j is an int

//depends on precision of constant
// i := 42           // int
// f := 3.142        // float64
// g := 0.867 + 0.5i // complex128

// consts cannot be declared with the := syntax

//an untyped constant takes the type needed by its context

// sum := 1
// 	for ; sum < 1000; {
// 		sum += sum
// 	}

// if you omit the loop condition it will loop forever

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %d: z = %f\n", i+1, z)
	}
	return z
}

func SqrtBetter(x float64) float64 {
	z := 1.0
	previous := 0.0
	for math.Abs(z-previous) > 1e-10 { //checks for absolute difference - if it is less than 1e-10 threshold, stops iteration
		previous = z
		z -= (z*z - x) / (2 * z)
		//fmt.Printf("Iteration %d: z = %f\n", i+1, z)
	}
	return z
}

// function for using slices
// func Pic(dx, dy int) [][]uint8 {
// 	result := make([][]uint8, dy) //make slice of length dy

// 	//populate rows
// 	for y := 0; y < dy; y++ {
// 		row := make([]uint8, dx)
// 		for x := 0; x < dx; x++ {
// 			row[x] = uint8(x ^ y)
// 		}
// 		result[y] = row
// 	}
// 	return result
// }

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(SqrtBetter(2))
	//pic.Show(Pic)
}

// switch os := runtime.GOOS; os {	//import runtime for this
// 	case "darwin":
// 		fmt.Println("macOS.")
// 	case "linux":
// 		fmt.Println("Linux.")
// 	default:
// 		// freebsd, openbsd,
// 		// plan9, windows...
// 		fmt.Printf("%s.\n", os

// switch without a condition is the same as switch true
// t := time.Now()
// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("Good morning!")
// 	case t.Hour() < 17:
// 		fmt.Println("Good afternoon.")
// 	default:
// 		fmt.Println("Good evening.")
// 	}

// defer fmt.Println("world")	//defers execution until surrounding function returns - so hello world is returned
// 	fmt.Println("hello")
// defer func calls are pushed to a stack - so are returned in LiFo order

// *T is a pnter to a T value
// & operator generates a pointer to its operand
// fmt.Println(*p) // read i through the pointer p
// *p = 21         // set i through the pointer p
// This is known as "dereferencing" or "indirecting".

// struct is a collection of fields

// struct literals - denotes a newly allocated struct value by listing the values of its fields
// type Vertex struct {
// 	X, Y int
// }
// var (
// 	v1 = Vertex{1, 2}  // has type Vertex
// 	v2 = Vertex{X: 1}  // Y:0 is implicit
// 	v3 = Vertex{}      // X:0 and Y:0
// 	p  = &Vertex{1, 2} // has type *Vertex
// )
// func main() {
// 	fmt.Println(v1, p, v2, v3)
// }

// names := [4]string{
// 		"John",
// 		"Paul",
// 		"George",
// 		"Ringo",
// 	}

// slice does not store any data, it just describes a section of an underlying array
// changing slice elements modifies the corresponding array elements
// slice literal is an array literal without the length
// slice has length and capacity - len(s) and cap(s)
// zero value of a slice is nil

// slices of slices
// board := [][]string{
// 		[]string{"_", "_", "_"},
// 		[]string{"_", "_", "_"},
// 		[]string{"_", "_", "_"},
// 	}
// board[0][0] = "X"
// 	board[2][2] = "O"

// appending to a slice
// func append(s []T, vs ...T) []T
// The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.
// The slice grows as needed.
// s = append(s, 1)
// printSlice(s)
// // We can add more than one element at a time.
// s = append(s, 2, 3, 4)
// printSlice(s)

// when using range for loops - _ skips the index/value
// if you only want the index, omit the second variable i.e for i := range pow

// map literals are like struct literals - but require keys
// var m = map[string]Vertex{
// 	"Bell Labs": Vertex{
// 		40.68433, -74.39967,
// 	},
// 	"Google": Vertex{
// 		37.42202, -122.08408,
// 	},
// }
// if the top-level type is just a type name, you can omit it from the elements of the literal

// Insert or update an element in map m:
// m[key] = elem
// Retrieve an element:
// elem = m[key]
// Delete an element:
// delete(m, key)
// Test that a key is present with a two-value assignment:
// elem, ok = m[key]
// If key is in m, ok is true. If not, ok is false.
// If key is not in the map, then elem is the zero value for the map's element type.
