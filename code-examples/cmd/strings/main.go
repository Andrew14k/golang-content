package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// // both return Hello World!
	// fmt.Println("\x48\x65\x6C\x6C\x6F\x20\x57\x6F\x72\x6C\x64\x21")
	// message := []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x20, 0x57, 0x6F, 0x72, 0x6C, 0x64, 0x21} //0x = hex code ^5,,6C etc = hex symbol
	// fmt.Println(string(message))

	// // var s1 string
	// // s1 = "Hello"
	// // var s2 string = "World!"
	// // s3 := s1 + " " + s2
	// //fmt.Println(s3)
	// //s4 := "Line1\n\tIndented Line2"	//indents

	// s6 := "String as a slice of bytes"
	// fmt.Println(string(s6[7:16])) //this notation results in a slice
	// fmt.Println(string(s6[:9]))   //from start to 9
	// fmt.Println(string(s6[7:]))   //from 7 to end

	// fmt.Println(strings.Index(s6, "s"))	//first index of s in string

	s := "    Go list fundamentals!     "

	s1 := strings.TrimSpace(s)
	s2 := strings.ToUpper(s1)
	s3 := strings.ReplaceAll(s2, "L", "*")
	s4 := strings.ReplaceAll(s3, "T", "*")
	s4 = s4[:strings.Index(s4, "N")+1]
	fmt.Println(s4)

	// strconv = parseInt, parseFloat, Atoi etc
	numString := "1298"
	if n, err := strconv.ParseInt(numString, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", n, n) //captial T for type
	}
}
