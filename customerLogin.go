package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func custoemerLogin() {
	for {
		phonenumber := GetInput("Enter phone number:")
		if true == back(phonenumber, 0) {
			break
		}
		password := GetInput("Enter password:")
		if true == back(password, 0) {
			break
		}
		for i, c := range customer {
			if c.phonenumber == phonenumber || c.password == password {

				choise := GetMenu("Customer panel", []string{"Change username and password", "change PhoneNumber", "My books", "New book"})
				if back(" ", choise) == true {
					break
				}
				switch choise {
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
				changeloop:
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
									break changeloop
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
				addLoop:
					for {
						if c.debt < 100000 {

							for _, b := range book {
								if b.quantity > 0 {
									fmt.Printf("name : %s .:*:. price : %d .:*:. ID : %s \n", b.name, b.price, b.ID)
								} else {
									fmt.Println("name :", b.name, " out of stock ")
								}
							}
							var ID []string

							for {
								fmt.Println("---------------------\nEnter IDs:")
								scanner.Scan()
								id := scanner.Text()
								if back(id, 0) == true {
									break addLoop
								}

								IDint, err := strconv.Atoi(id)

								if err != nil {
									fmt.Println("err while converting id")
									continue
								}

								if IDint < 1 || IDint > len(book) {
									fmt.Println("id does not exist")
									continue
								}

								ID = append(ID, id)

							}
							fileSelf, err := os.OpenFile("src/"+c.name, os.O_CREATE|os.O_WRONLY, 0644)
							if err != nil {
								fmt.Println("err while openign fileSelf", err)
								break addLoop
							}
							fileCustomer, err := os.OpenFile("src/customer.txt", os.O_WRONLY, 0644)

							if err != nil {
								fmt.Println("err while opening customer ", err)
								closer([]*os.File{fileSelf})
								break addLoop
							}
							fileBook, err := os.OpenFile("src/book.txt", os.O_WRONLY, 0644)
							if err != nil {
								fmt.Println("err while openign bookFile", err)
								closer([]*os.File{fileSelf, fileCustomer})
								break addLoop
							}
							_, err = fileCustomer.Seek(int64(i*41)+31, io.SeekStart)
							if err != nil {
								fmt.Println("err while writing debt", err)
								closer([]*os.File{fileBook, fileSelf, fileCustomer})
								break addLoop
							}
							for _, id := range ID {
								i, err := strconv.Atoi(id)
								i = i - 1
								if err != nil {
									fmt.Println("err while convertin", err)
									closer([]*os.File{fileBook, fileSelf, fileCustomer})
									break addLoop

								}
								_, err = fileBook.Seek(int64((i*33)+23), io.SeekStart)
								if err != nil {
									fmt.Println("err while seeking book", err)

									break
								}
								_, err = fileSelf.Seek(0, io.SeekEnd)
								if err != nil {
									fmt.Println("err while seeking book", err)

									break
								}
								var b string
								if book[i].quantity-1 > 0 {
									b := strconv.Itoa(book[i].quantity - 1)
									b = FillDot(b, 8)
								} else {
									fmt.Println("out of stock")
								}
								fileCustomer.Seek(int64((i*41)+31), io.SeekStart)
								fileBook.WriteString(b)
								fileSelf.WriteString(id + "\r\n")

							}
							closer([]*os.File{fileSelf, fileBook, fileCustomer})

						} else {
							fmt.Println("you need to pay debt first")
							break
						}
					}
				}
			}
		}
	}
}
