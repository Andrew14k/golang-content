// go functions may be closures
// a closure is a function value that references variables from outside its body
// the function may access and assign to the referenced variables. in the sense the function is bound to the variables
// For example, the adder function returns a closure. Each closure is bound to its own sum variable.
// func adder() func(int) int {
// 	sum := 0
// 	return func(x int) int {
// 		sum += x
// 		return sum
// 	}
// }

// func main() {
// 	pos, neg := adder(), adder()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(
// 			pos(i),
// 			neg(-2*i),
// 		)
// 	}
// }

// Go doesnt have classes, but can define methods on types - method is function with a special receiver class
// type Vertex struct {
// 	X, Y float64
// }
// func (v Vertex) Abs() float64 {	//Abs method has receiver type Vertex named v
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }
// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(v.Abs())
// }

// can declare methods with pointer receivers
// functions with a pointer argument must take a pointer
// ScaleFunc(v, 5)  // Compile error!
// ScaleFunc(&v, 5) // OK
// methods with pointer receivers take either a value or a pointer as the receiver when they are called
// v.Scale(5)  // OK
// p := &v
// p.Scale(10) // OK

// equivalent thing happens in the reverse direction
// Functions that take a value argument must take a value of that specific type:
// var v Vertex
// fmt.Println(AbsFunc(v))  // OK
// fmt.Println(AbsFunc(&v)) // Compile error!
// while methods with value receivers take either a value or a pointer as the receiver when they are called:

// var v Vertex
// fmt.Println(v.Abs()) // OK
// p := &v
// fmt.Println(p.Abs()) // OK
// In this case, the method call p.Abs() is interpreted as (*p).Abs().

// reasons to use a pointer receiver = so ethod can modify the value that the receiver points to, and avoid copying the value on each method call
// generally, all methods on a given type should have either value or pointer receivers, but not a mixture of both

// INTERFACES
// an interface type is defined as a set of method signatures, and a value of interface type can hold any value that implements those methods
// a type implements an interface by implementing its methods
// implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement

// interface values can be thought of as a tuple of a value and a concrete type
// calling a method on an interface value executes the method of the same name on its underlying type
// type I interface {
// 	M()
// }
// type T struct {
// 	S string
// }
// func (t *T) M() {
// 	fmt.Println(t.S)
// }
// type F float64

// func (f F) M() {
// 	fmt.Println(f)
// }
// func main() {
// 	var i I

// 	i = &T{"Hello"}
// 	describe(i)
// 	i.M()

// 	i = F(math.Pi)
// 	describe(i)
// 	i.M()
// }
// func describe(i I) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// common in go to write methods that handle being called with a nil receiver -
// an interface value that holds a nil concrete value is itself non-nil

// TYPE ASSERTIONS

// type assertion provides access to an interface values underlying concrete value
// t := i.(T)
// This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.

// t, ok := i.(T) //tests an interface value holds a specific type
// var i interface{} = "hello"
// 	s, ok := i.(string)
// 	fmt.Println(s, ok)
// 	f, ok := i.(float64)
// 	fmt.Println(f, ok)

// STRINGER INTERFACE - type that can describe itself as a string
// type Person struct {
// 	Name string
// 	Age  int
// }
// func (p Person) String() string {
// 	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
// }
// func main() {
// 	a := Person{"Arthur Dent", 42}
// 	z := Person{"Zaphod Beeblebrox", 9001}
// 	fmt.Println(a, z)
// }

// ERRORS

// type error interface {
//     Error() string
// }

// func (e *MyError) Error() string {
// 	return fmt.Sprintf("at %v, %s",
// 		e.When, e.What)
// }

// i, err := strconv.Atoi("42")
// if err != nil {
//     fmt.Printf("couldn't convert number: %v\n", err)
//     return
// }
// fmt.Println("Converted integer:", i)

// READERS
// read method populates the given byte slice with data and returns the number of bytes populated and an error value
// returns io.EOF error when the stream ends
// code below creates a string.Reader and consumes its output 8 bytes at a time

// package main

// import (
// 	"fmt"
// 	"io"
// 	"strings"
// )

// func main() {
// 	r := strings.NewReader("Hello, Reader!")

// 	b := make([]byte, 8)
// 	for {
// 		n, err := r.Read(b)
// 		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
// 		fmt.Printf("b[:n] = %q\n", b[:n])
// 		if err == io.EOF {
// 			break
// 		}
// 	}
// }

// IMAGES
// package image

// type Image interface {
//     ColorModel() color.Model
//     Bounds() Rectangle
//     At(x, y int) color.Color
// }
// these two are independent of each other
// func main() {
// 	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
// 	fmt.Println(m.Bounds())
// 	fmt.Println(m.At(0, 0).RGBA())
// }