package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func customerLogin() {
Exit:
	for {
		phonenumber := GetInput("Enter phone number:")
		if true == back(phonenumber, 0) {
			break
		}
		password := GetInput("Enter password:")
		if true == back(password, 0) {
			break
		}
		for {
			for i, c := range customer {
				if c.phonenumber == phonenumber || c.password == password {

					choice := GetMenu("Customer panel", []string{"Change username and password", "change PhoneNumber", "My books", "New book", "Pay debt"})
					if back(" ", choice) == true {
						break Exit
					}
					switch choice {
					case 1:

						for {
							fmt.Printf("your user name : %s your password : %s \n Enter new username and password \n", c.name, c.password)
							scanner.Scan()
							newUsername := strings.TrimSpace(scanner.Text())
							if true == back(newUsername, 0) {
								break
							}
							scanner.Scan()
							newPassword := strings.TrimSpace(scanner.Text())
							if true == back(newPassword, 0) {
								break
							}
							for range customer {

								if _, err := os.Stat(c.name); err == nil {

									err := os.Rename(c.name, newUsername)
									if err != nil {
										fmt.Println("err while renamign file", err)
									}
								}

								c.name = newUsername
								c.password = newPassword
								newUsername = FillDot(c.name, 10)
								newPassword = FillDot(c.password, 10)

								file, err := os.OpenFile("src/customer.txt", os.O_WRONLY, 0644)
								if err != nil {
									fmt.Println("err while opening file", err)

								}
								file.Seek(int64((i*41)+11), io.SeekStart)
								file.WriteString(newUsername)
								file.WriteString(newPassword)

								closer([]*os.File{file})
								break
							}
						}
					case 2:
					changeLoop:
						for {
							fmt.Println("your phonenumber :", c.phonenumber, "\n Enter new phonenumber :")
							scanner.Scan()
							input := scanner.Text()
							back(input, 0)
							if len(input) == 11 {
								c.phonenumber = input
								file, err := os.OpenFile("src/customer.txt", os.O_WRONLY, 0644)
								if err != nil {
									fmt.Println("err while opening file", err)
								}
								for _, x := range customer {
									if checkName(input, x.phonenumber) == false {
										fmt.Println("This phonenumber already exist")
										break changeLoop
									}
								}
								file.Seek(int64(i*41), io.SeekStart)
								file.WriteString(input)
								fmt.Println("your new phonenumber :", c.phonenumber)

								closer([]*os.File{file})
								break
							} else {
								fmt.Println("phonenumber most be 11c")
							}

						}
					case 3:

					case 4:
					Loop:
						for {
							if c.debt > 1000000 {
								fmt.Println("you should pay your debt first ")
								break
							}
							for _, b := range book {
								fmt.Println(b.name, b.price, b.ID)
							}
							fileSelf, err := os.OpenFile("src/"+c.name, os.O_WRONLY|os.O_CREATE, 0644)
							if err != nil {
								fmt.Println("err while opening fileSelf", err)
								break

							}
							fileBook, err := os.OpenFile("src/book.txt", os.O_WRONLY, 0644)
							if err != nil {
								fmt.Println("err while opening fileBook", err)
								closer([]*os.File{fileSelf})
								break
							}
							fileCustomer, err := os.OpenFile("src/customer.txt", os.O_WRONLY, 0644)
							if err != nil {
								fmt.Println("err while opening customer.txt", err)
								closer([]*os.File{fileSelf, fileBook})
								break

							}
							for {
								input := GetInput("Enter ids: ")
								if back(input, 0) {
									break Loop
								}
								found := false
								for indexB, b := range book {
									if false == checkName(input, b.ID) {
										found = true
										_, err = fileBook.Seek(int64(indexB*33)+18, io.SeekStart)
										if err != nil {
											fmt.Println("err while seeking book", err)
											break
										}
										_, err = fileCustomer.Seek(int64(i*41)+31, io.SeekStart)
										if err != nil {
											fmt.Println("err while seeking customer", err)
											break
										}
										_, err := fileSelf.Seek(0, io.SeekEnd)
										if err != nil {

											fmt.Println("err while seeking fileSelf", err)
											break
										}
										book[indexB].quantity = book[indexB].quantity - 1
										customer[i].debt = customer[i].debt + book[indexB].price

										fileSelf.WriteString(input + "\r\n")
										fileBook.WriteString(FillDot(strconv.Itoa(book[indexB].quantity), 5))
										fileCustomer.WriteString(FillDot(strconv.Itoa(customer[i].debt), 8))
									}
								}
								if found != true {
									fmt.Println("unknown id")
								}
							}
							closer([]*os.File{fileSelf, fileBook, fileCustomer})
						}
					case 5:
					}

				}
			}
		}
	}
}
