package main

import (
	"fmt"
	"log"

	"example.com/dockerexample/internal/db"
	"example.com/dockerexample/internal/models"
	"example.com/dockerexample/internal/repository"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	if err := db.New(); err != nil {
		log.Fatalf("%v", err)
	}
	defer db.Close() //code after this would be the assumption here at this point

	GetCustomers()

	GetCustomerById()

	CreateCustomer()

	UpdateCustomerContact()

	DeleteCustomer()

}

func UpdateCustomerContact() {
	beforeUpdate, afterUpdate, err := repository.UpdateCustomerContact(94, "Clara Oswald")
	if err != nil {
		log.Fatalf("failed to update customer contact name: %w", err)
	}
	fmt.Println("Before update: ")
	for _, c := range beforeUpdate {
		fmt.Printf("%d, %s, %s\n", c.CustomerID, c.CustomerName, c.ContactName)
	}
	fmt.Println("After update: ")
	for _, c := range afterUpdate {
		fmt.Printf("%d, %s, %s\n", c.CustomerID, c.CustomerName, c.ContactName)
	}
}

func DeleteCustomer() {
	if err := repository.DeleteCustomer(98); err != nil {
		log.Fatalf("failed to delete customer: %v", err)
	}
	fmt.Printf("Deleted customer with ID %d\n", 98)
}

func CreateCustomer() {
	newCustomer := models.Customer{
		CustomerName: "Gallifrey Co",
		ContactName:  "Omega",
		Address:      "Kastaboros",
		City:         "The Citadel",
		PostalCode:   "234124",
		Country:      "NA",
	}
	customerCreated, err := repository.CreateCustomer(newCustomer)
	if err != nil {
		log.Fatalf("failed to create customer: %v", err)
	}
	fmt.Printf("Created customer: %+v", customerCreated)
}

func GetCustomerById() {
	customerByID, err := repository.GetCustomerById(77)
	if err != nil {
		log.Fatalf("failed to get customer by ID: %v", err)
	}
	fmt.Printf("Customer with ID %d: %v\n\n", 77, customerByID)
}

func GetCustomers() {
	// call GetCustomers
	customers, err := repository.GetCustomers()
	if err != nil {
		log.Fatalf("failed to get customers: %v", err)
	}

	// print customer data
	for _, c := range customers {
		fmt.Printf("%d, \t%s,\t %s, \t%s, \t%s \n", c.CustomerID, c.CustomerName, c.ContactName, c.City, c.Country)
	}
}

// func GetCustomers() {
// 	rows, err := db.Conn.Query("SELECT CustomerId, CustomerName, City FROM Customers")
// 	if err != nil {
// 		log.Fatalf("failed to run query on customer table: %v", err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var id int
// 		var name string
// 		var city string

// 		err := rows.Scan(&id, &name, &city)
// 		if err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			continue
// 		}
// 		fmt.Printf("ID: %d,\tName: %s,\tCity: %s\n", id, name, city)
// 	}
// }

// FINDING FOREIGN KEY NAME
/* SELECT CONSTRAINT_NAME
FROM information_schema.KEY_COLUMN_USAGE
WHERE TABLE_NAME = 'Orders' AND COLUMN_NAME = 'CustomerID' AND CONSTRAINT_SCHEMA = 'your_database_name'; */
//alternate is to use DELETE CASCADE???
