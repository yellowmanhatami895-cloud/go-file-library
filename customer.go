package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Customer struct {
	PhoneNumber string
	Name        string
	Password    string
	Debt        int
}

func readCustomerFromFile(file *os.File) (Customer, error) {
	line, err := readLineFromFile(file)
	if err != nil {
		return Customer{}, err
	}
	var customer Customer
	fields := strings.Split(line, "|")
	if len(fields) < 4 {
		return Customer{}, fmt.Errorf("invalid record")
	}
	n := fields[1]
	pas := fields[2]
	i, err := strconv.Atoi(fields[3])
	if err != nil {
		fmt.Println("err conv inventory(customer) : ", err)
	}
	customer = Customer{
		PhoneNumber: fields[0],
		Name:        n,
		Password:    pas,
		Debt:        i,
	}
	return customer, nil

}
func readCustomer() {
	file, err := os.OpenFile("src/customer.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("err while opening customer.txt")
		fmt.Println(err)
	}
	defer file.Close()
	for {
		customer, err := readCustomerFromFile(file)
		if err != nil {
			continue
		}
		customers = append(customers, customer)
	}
}
