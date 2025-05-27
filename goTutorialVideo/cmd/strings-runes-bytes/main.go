package main

import (
	"fmt"
	"strings"
)

// go represents strings with utf8 encoding
// rune is an alias for n32
func main() {
	var myString = []rune("résumé")
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed) //index returned in the bytearray, will skip byte for special charcters like the accent e
	for i, v := range myString {
		fmt.Println(i, v)
	}
	// strings are immutable in go, cannot be modified once created
	var strSlice = []string{"h", "e", "y"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf(catStr)
	// same but with string package
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr2 = strBuilder.String()
	fmt.Println(catStr2)

}
