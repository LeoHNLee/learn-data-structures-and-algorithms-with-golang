package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templateHtml = template.Must(template.ParseGlob("templates/*"))

// Create method
func Create(writer http.ResponseWriter, request *http.Request) {
	templateHtml.ExecuteTemplate(writer, "Create", nil)
}

// Insert
func Insert(writer http.ResponseWriter, request *http.Request) {
	var customer Customer
	customer.CustomerName = request.FormValue("customername")
	customer.SSN = request.FormValue("ssn")
	InsertCustomer(customer)
	customers := GetCustomers()
	templateHtml.ExecuteTemplate(writer, "Home", customers)
}

// Alter
func Alter(writer http.ResponseWriter, request *http.Request) {
	var customer Customer
	var customerId int
	customerIdStr := request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	customer.CustomerId = customerId
	customer.CustomerName = request.FormValue("customername")
	customer.SSN = request.FormValue("ssn")
	UpdateCustomer(customer)
	customers := GetCustomers()
	templateHtml.ExecuteTemplate(writer, "Home", customers)
}

// Update
func Update(writer http.ResponseWriter, request *http.Request) {
	var customerId int
	customerIdStr := request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", customerId)
	customer := GetCustomerById(customerId)
	templateHtml.ExecuteTemplate(writer, "Update", customer)
}

// Delete
func Delete(writer http.ResponseWriter, request *http.Request) {
	var customerId int
	customerIdStr := request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	customer := GetCustomerById(customerId)
	DeleteCustomer(customer)
	customers := GetCustomers()
	templateHtml.ExecuteTemplate(writer, "Home", customers)
}

// View
func View(writer http.ResponseWriter, request *http.Request) {
	var customerId int
	customerIdStr := request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	customer := GetCustomerById(customerId)
	fmt.Println(customer)
	customers := []Customer{customer}
	// customers = append(customers, customer)
	templateHtml.ExecuteTemplate(writer, "View", customers)
}

// Home
func Home(writer http.ResponseWriter, reader *http.Request) {
	customers := GetCustomers()
	log.Println(customers)
	templateHtml.ExecuteTemplate(writer, "View", customers)
}

// main method
func main() {
	log.Println("Server started on: http://localhost:8000")
	http.HandleFunc("/", Home)
	http.HandleFunc("/alter", Alter)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/view", View)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8000", nil)
}
