package main

import "fmt"

func main() {
	bankApp()
}

func bankApp() {
	var balance float64 = 0
	var choice int
	var amount float64
	for {
		fmt.Println(" ")
		fmt.Println("ATM Options:")
		fmt.Println("1. Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)
		fmt.Println(" ")

		switch choice {
		case 1:
			fmt.Printf("Balance is %f", balance)
			fmt.Println(" ")
		case 2:
			fmt.Print("Enter amount of deposit: ")
			fmt.Scanln(&amount)
			balance += amount
			fmt.Printf("New balance is %f", balance)
			fmt.Println(" ")
		case 3:
			fmt.Print("Enter amount to refund: ")
			fmt.Scanln(&amount)
			balance -= amount
			fmt.Printf("New balance is %f", balance)
			fmt.Println(" ")
		case 4:
			fmt.Printf("Exiting. balance is %f", balance)
			fmt.Println(" ")
			return
		default:
			fmt.Println("Invalid input. Please select 1-4.")
			fmt.Println(" ")
		}
	}
}

//int on its own (instead of declaring int32 or int64, will change based on its system)
