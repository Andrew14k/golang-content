// []Arrays = fixed length, same type, indexable, contiguous in memory
package main

import "fmt"

func main() {
	var intArr [3]int32 //array of length 3 with [0,0,0] as elements
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3]) //items 1 to 2
	//fmt.Println(&inArr[0]) // tells you memory location of intArr[0]

	//immediately initialise array:
	// var intArr [3]int32 = [3]int32{1,2,3}
	// intArr := [...]int32{1,2,3}	//array size inferred by compiler, can but still put 3 in []

	// SLICES - wrap arrays to give more genral, powerful and convenient interface to sequences of data
	//slices can omit length value to have slice
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7) //basically creates a new array, with all elements plus extra capacity
	fmt.Println(intSlice)          //[4,5,6,7,*,*]
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...) //[4,5,6,7,8,9]

	//alternate slice
	var intSlice3 []int32 = make([]int32, 3, 8) //3 is length of slice, optionally can specify capacity of slice (8 here)
	fmt.Println(intSlice3)

	//MAPS - set of key value pairs MAPS ALWAYS RETURN SOMETHING EVEN IF KEY DOESNT EXIST
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["jason"]
	//delete(myMap2, "Adam")	//no return value given
	if ok { //checks for if key in map
		fmt.Printf("age is %v", age)
	} else {
		fmt.Println("invalid name")
	}

	// LOOPS
	// use range to iterate over map, array or slice
	for name, age := range myMap2 {
		fmt.Printf("Name %f, Age: %v \n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("index: %v, Value: %v \n", i, v)
	}
	//GO DOESNT HAVE WHILE LOOPS - achieved with for loops
	var i int = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}
	//OR
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
