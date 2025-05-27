package main

import (
	"fmt"
	"unicode/utf8"
) //have to use packages that are imported

// go build main.go creates a binary file named main to run code
func main() {
	var intNum int = 32767 //can use int16, int32, int64 for specific memory allocations
	fmt.Println(intNum)

	var floatNum float32 = 10.1 //must specify 32 or 64 bit for floats
	fmt.Println(floatNum)

	var result float32 = floatNum + float32(intNum)
	fmt.Println(result)

	var int1 int = 3
	var int2 int = 4
	fmt.Println(int1 / int2)

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	fmt.Println(len("test"))                    //4 - number of bytes, not actually number of characters
	fmt.Println(utf8.RuneCountInString("word")) //works for special characters also

	var myRune = 'a'
	fmt.Println(myRune)

	//default values for numberical types is 0, empty string for strings and chars

	myVar := "text"
	fmt.Println(myVar)
	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	//booleans are normal
}
