package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 { //asigning function/(method now) to gas engine type
	return e.gallons * e.mpg
}
func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

// allows for more general functions
type engine interface {
	milesLeft() uint8
}

// func canMakeIt(e gasEngine, miles uint8) {	//not general - use interfaces to allow for any engine not just gasEngine
// 	if miles <= e.milesLeft() {
// 		fmt.Println("You can make it")
// 	} else {
// 		fmt.Println("fuel up mate")
// 	}
// }
func canMakeIt(e engine, miles uint8) { //not general - use interfaces to allow for any engine not just gasEngine
	if miles <= e.milesLeft() {
		fmt.Println("You can make it")
	} else {
		fmt.Println("fuel up mate")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}} //can replace gasEngine with electricEngine and canMakeIt will still work
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)
	fmt.Printf("Miles left in tank: %v", myEngine.milesLeft())
	canMakeIt(myEngine, 50)
}

//anonymous structs - define and initialise in same place - i.e. in main
