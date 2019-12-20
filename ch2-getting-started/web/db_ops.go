package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Customer Class
type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

// GetConnection method which returns sql.DB
func GetConnection() *sql.DB {
	databaseDriver := "mysql"
	databaseUser := "root"
	databasePass := "Ms8pw8e85344!"
	databaseName := "tcp(127.0.0.1:3306)/golang"

	// 실제 커넥션이 이뤄지지 않는다.
	// 커넥션 정보조차 체크하지 않으며
	// 실제 커넥션은 query 단계에서 이뤄진다.
	database, error := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@"+databaseName)
	if error != nil {
		panic(error.Error())
	}
	return database
}

// GetCustomers method returns Customer Array
func GetCustomers() []Customer {
	database := GetConnection()

	rows, error := database.Query("SELECT * FROM Customer ORDER BY CustomerId DESC")
	if error != nil {
		panic(error.Error())
	}

	customer := Customer{}
	customers := []Customer{}
	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		error = rows.Scan(&customerId, &customerName, &ssn)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)
	}
	defer database.Close()
	return customers
}

// InsertCustomer method with parameter customer
func InsertCustomer(customer Customer) {
	database := GetConnection()

	insert, error := database.Prepare("INSERT INTO CUSTOMER(CustomerName, SSN) VALUES(?,?)")
	if error != nil {
		panic(error.Error())
	}
	insert.Exec(customer.CustomerName, customer.SSN)
	defer database.Close()
}

// UpdateCustomer method with parameter customer
func UpdateCustomer(customer Customer) {
	database := GetConnection()
	update, err := database.Prepare("update customer set CustomerName=?, SSN=? where CustomerId=?")
	if err != nil {
		panic(err.Error())
	}
	update.Exec(customer.CustomerName, customer.SSN, customer.CustomerId)
	defer database.Close()
}

// GetCustomerById with parameter customerId returns Customer
func GetCustomerById(customerId int) Customer {
	database := GetConnection()
	rows, err := database.Query("select * from Customer where CustomerId=?", customerId)
	if err != nil {
		panic(err.Error())
	}
	customer := Customer{}
	for rows.Next() {
		var customerId int
		var customerName string
		var SSN string
		err := rows.Scan(&customerId, &customerName, &SSN)
		if err != nil {
			panic(err.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = SSN
	}
	defer database.Close()
	return customer
}

// DeleteCustomer method with parameter customer
func DeleteCustomer(customer Customer) {
	database := GetConnection()
	delete, err := database.Prepare("delete from Customer where CustomerId =?")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(customer.CustomerId)
	defer database.Close()
}
