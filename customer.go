package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Customer struct {
	phonenumber string
	name        string
	password    string
	debt        int
}

func read_customer() {
	file, err := os.OpenFile("src/customer.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("err while opening customer.txt")
		fmt.Println(err)
	}
	defer file.Close()
	Line := 0
	for {
		file.Seek(int64(Line*41), io.SeekStart)
		by := make([]byte, 39)
		_, err = file.Read(by)
		if err == io.EOF {
			fmt.Println("EOF customer")
		}
		if err != nil {
			break
		}
		n := strings.Trim(string(by[11:21]), ".")
		pas := strings.Trim(string(by[21:31]), ".")
		i, err := strconv.Atoi(strings.Trim(string(by[31:]), "."))
		if err != nil {
			fmt.Println("err conv inventory(customer) : ", err)
		}
		customer = append(customer, Customer{
			phonenumber: string(by[:11]),
			name:        n,
			password:    pas,
			debt:        i,
		})

		Line++
	}

}
