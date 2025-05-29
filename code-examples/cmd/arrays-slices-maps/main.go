package main

import (
	"fmt"
	"reflect"
)

func main() {
	//ARRAYS - fixed size - value type
	fmt.Println("ARRAYS")
	x := [3]int{1, 2, 3}
	y := x //y now has copy of what x holds - copy of the value
	y[0] = 5
	fmt.Println(x[0]) //1
	fmt.Println(y[0]) //5
	fmt.Println(" ")

	//SLICES - reference type - a descriptor of an underlying array
	fmt.Println("SLICES")
	a := []int{1, 2, 3}
	b := a
	b[0] = 5          //changes reflected in both arrays - change affects both
	fmt.Println(a[0]) //5
	fmt.Println(b[0]) //5
	fmt.Println(" ")

	//LENGTH AND CAPACITY
	fmt.Println("LENGTH AND CAPACITY WITH SLICES")
	fmt.Println(len(a), cap(a)) // 3,3
	a = append(a, 4)
	fmt.Println(len(a), cap(a)) // 3,6 - when slice is full, any appendages will double the capacity so if len and cap are the same, an append would make len+1 and cap*2
	fmt.Println(" ")

	// s := make([]int, 3, 10)	//integer slice, len, capacity
	// fmt.Println(s)	//0,0,0

	fmt.Println("POINTERS FOR SLICES")
	ra := reflect.ValueOf(a)
	rb := reflect.ValueOf(b)
	pa := ra.Pointer()
	pb := rb.Pointer()
	fmt.Printf("Pointer held by slice a: %#x - Pointer held by slice b: %#x", pa, pb)
	fmt.Println(" ")
	fmt.Println(" ")

	fmt.Println("ARRAYS WITH SLICES")
	arr := [5]int{1, 2, 3, 4, 5}
	sliceOfArray := arr[:3] //description on the array value

	for _, val := range sliceOfArray { // prints 1,2,3 on one line after the other
		fmt.Println(val)
	}
	sliceOfArray[0] = 2                             //much easier and cleaner than using the append function
	fmt.Println("Value of element in arr:", arr[0]) //slice points to eact array that array is pointing to - not a copy
	fmt.Println(" ")

	fmt.Println("MAPS") //associative array - can specify keys, maps allow for fast lookups, insertions, deletions
	//maps are reference types, but unlike slices, they cannot guarantee the order of elements in a map
	// slices maps and functiosn cannot be used as keys

	//make syntax for maps - or map literal (shown below)
	daysInMonth := map[string]int{"January": 31, "February": 28, "March": 31} //map[key]value{k1:v1, k2,v2}
	for month, days := range daysInMonth {
		fmt.Printf("There are %d days in %s\n", days, month)
	}

	// if you print mutliple times, you will see apple and banana can be output in different orders
	//example with make
	produce := make(map[string]float32)
	produce["apple"] = 1.2
	produce["banana"] = 0.5

	for item, price := range produce {
		fmt.Printf("Item: %s - £%.2f\n", item, price)
	}
	//make(map[key]value)
	//m[key] := value	//add or update value
	//val := m[key] //retireves value

	// check if a val exists - ok is a bool to check if exists
	// 	if val, ok := m[key]; ok {
	//   // execute if key found
	// } else {
	//   // execute if key isn't found
	// }

	//delete(m, key)
	//len(m)
	//for k,v := range m
	fmt.Println(" ")
}
