package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"example.com/dockerexample/internal/db"
	"example.com/dockerexample/internal/models"
	"github.com/go-sql-driver/mysql"
)

func GetCustomers() ([]*models.Customer, error) {
	customers := []*models.Customer{}
	rows, err := db.Conn.Query("SELECT * FROM Customers")
	if err != nil {
		return nil, fmt.Errorf("failed to run query on customer table: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var name string
		var contact string
		var addy string
		var city string
		var postcode string
		var country string

		if err := rows.Scan(&id, &name, &contact, &addy, &city, &postcode, &country); err != nil {
			log.Printf("failed to scan row: %v", err)
			continue
		}

		customer := &models.Customer{
			CustomerID:   id,
			CustomerName: name,
			ContactName:  contact,
			Address:      addy,
			City:         city,
			PostalCode:   postcode,
			Country:      country,
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func GetCustomerById(id int64) (*models.Customer, error) {
	c := &models.Customer{}
	row := db.Conn.QueryRow("SELECT * FROM Customers WHERE CustomerId = ?", id)
	err := row.Scan(
		&c.CustomerID,
		&c.CustomerName,
		&c.ContactName,
		&c.Address,
		&c.City,
		&c.PostalCode,
		&c.Country)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("no customer found: %w", err)
	}
	return c, nil
}

func CreateCustomer(c models.Customer) (*models.Customer, error) {
	result, err := db.Conn.Exec("INSERT INTO Customers (Customername, ContactName, Address, City, PostalCode, Country) VALUES (?, ?, ?, ?, ?, ?)",
		c.CustomerName,
		c.ContactName,
		c.Address,
		c.City,
		c.PostalCode,
		c.Country,
	) //removed c.CustomerID
	if err != nil {
		return nil, fmt.Errorf("customer not saved: %w", err)
	}

	id, _ := result.LastInsertId()
	c.CustomerID = id
	return &c, nil
}

func DeleteCustomer(id int64) error {
	_, err := db.Conn.Exec("ALTER TABLE Orders DROP FOREIGN KEY Orders_ibfk_2")
	if err != nil {
		//if sql error - if foreign key already dropped
		var sqlErr *mysql.MySQLError //derference to get value of sql error
		if errors.As(err, &sqlErr) && sqlErr.Number == 1091 {
			log.Println("Foreign key already dropped, continuing...")
		} else {
			return fmt.Errorf("failed to drop foreign key constraints: %w", err)
		}
	}

	// foreign key constraint removed, so can delete from Customers now
	result, err := db.Conn.Exec("DELETE FROM Customers WHERE CustomerID = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete customer: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no customer found with ID %d", id)
	}
	return nil
}

func UpdateCustomerContact(id int64, newContact string) ([]*models.Customer, []*models.Customer, error) {
	beforeUpdate, err := GetCustomers()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get customers from db: %w", err)
	}

	_, err = db.Conn.Exec("UPDATE Customers SET ContactName = ? WHERE CustomerId = ?", newContact, id)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to update contact name: %w", err)
	}

	afterUpdate, err := GetCustomers()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get updated customers from db: %w", err)
	}
	return beforeUpdate, afterUpdate, nil
}
