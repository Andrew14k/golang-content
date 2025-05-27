// variables that store memory locations instead of values (ints, floats)
package main

import "fmt"

func main() {
	//var p *int32	//at this point, it has no address assigned to it
	// var p *int32 = new(int32)                  //points to memory address
	// fmt.Printf("value p points to is: %v", *p) //using *p is called dereferencing the pointer
	// var i int32                                //defaults to 0
	// //*p = 10                                    //set value of memory location p is pointing to to 10
	// p = &i //memory address of i

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nmemory location of thing1 is: %p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\nresult is: %v", result)
	fmt.Printf("\nvalue of thing1 is: %v", thing1)
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nmemory location of thing2 is: %p", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}

//these would be two different arrays unless the * and & pointers is used to ensure that they are stored in the same memory location
//pointers are useful when managing large parameters, avoiding wasting time and memory
//values of thing2 and thing1 change togetehr
