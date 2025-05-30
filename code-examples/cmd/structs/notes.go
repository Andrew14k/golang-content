package main

import "fmt"

// struts allow you to bundle data
// defined at package level - outside of any function
// can define anonymous strcuts within functions - useful for testing (table approach)

type Person struct { //capitalisation makes it exportable outside this package, so fields will need to be capitalised to also be exportable
	Name string //usually keep fields invisiable - not capitalised - access method usually used
	Age  int
}

type address struct {
	city    string
	country string
}

type person2 struct {
	name string
	age  int
	address
}

// METHODS
func (a address) getAddress() string { //deals with copy of the underlying data structure
	return fmt.Sprintf("%s, %s\n", a.city, a.country)
	// this is just a copy so if you made a a.city = "" in here it wouldnt be changed in main, so in theory everything we do in here is safe
	// if you do want to allow for changed would use func (a *address) - pointer to memory location so changes will take affect
}

func (a *address) changeCity(city string) { //deals with a reference to the underlying data structure - changes city
	a.city = city
}

func notes() {
	////////////////////////////////////////////////////////////////////////////////////////////
	//INTIAL STRUCTS WORK //////////////////
	// p := Person{ // always include commas even on the last elements
	// 	Name: "Alice",
	// 	Age:  22,
	// }

	// p2 := Person{"James", 32}

	// var p3 Person
	// p3.Name = "Dave"
	// p3.Age = 48

	// // below is same as var p4 *Person
	// p4 := new(Person) //using new keyword returns a pointer to a zero-initialised Person struct - holds pointer rather than the struct itself
	// p4.Name, p4.Age = "Steve", 39

	// p5 := &Person{ //p5 returns pointer to Person - same as p4
	// 	Name: "Sandra",
	// 	Age:  55,
	// }
	// fmt.Println(p, p2, p3, p4, p5)
	//////////////////////////////////////////////////////////////////////////////////////////////

	// go is a compositional language - doesnt support class-based inheritance
	// can do structs within structs instead

	p := person2{
		name: "andrew",
		age:  23,
	}
	p.city = "London"
	p.country = "UK"
	fmt.Printf("%s (%d) is from %s\n", p.name, p.age, p.getAddress())

	p.changeCity("Scunthorpe")
	fmt.Printf("%s (%d) is from %s\n", p.name, p.age, p.getAddress())
}
