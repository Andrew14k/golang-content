package main

import "fmt"

func main() {
	testPrecedence()
	calc()

}

func testPrecedence() {
	fmt.Println(10 + 3 - 2)               //returns 11
	fmt.Println((10 + 3) * 2)             //returns 26 - remove brackets this would return 16
	fmt.Println(4 > 2 || 5 < 1 && 3 == 3) //returns true
	fmt.Println(true || false && false)   //returns true
	fmt.Println(10+4 > 12 && 5 < 10)      //returns true
	// 10 /3 would return 3, need to clarify type as float64 or float32 to ensure decimal return

	var x int = 10
	var y float64 = 2.5
	fmt.Println(float64(x) + y)

	//cannot operate on mismatched types
	//have to ensure matched typed for relational operations also (>,<,==,...)
	//logical operators in Go = &&, || and !
}

func calc() { //currently for only integers
	var num1 int
	var op string
	var num2 int
	fmt.Print("Enter first number")
	fmt.Scanln(&num1)

	fmt.Print("Enter operator")
	fmt.Scanln(&op)

	fmt.Print("Enter second number")
	fmt.Scanln(&num2)

	switch op {
	case "+":
		fmt.Printf("Result: %d + %d = %d\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("Result: %d - %d = %d\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("Result: %d * %d = %d\n", num1, num2, num1*num2)
	case "/":
		fmt.Printf("Result: %d / %d = %d\n", num1, num2, num1/num2)
	case "%":
		fmt.Printf("Result: %d %% %d = %d\n", num1, num2, num1%num2)
	}
}
