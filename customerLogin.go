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
		phonenumber := GetInput("Enter phone number [0 for back]:")
		if phonenumber == "0" {
			break
		}
		password := GetInput("Enter password [0 for back]:")
		if password == "0" {
			break
		}
		for {
			for i, c := range customers {
				if c.PhoneNumber == phonenumber || c.Password == password {

					choice := GetMenu("Customer panel", []string{"Change username and password", "change PhoneNumber", "My books", "New book", "Pay debt"})
					if choice == 0 {
						break Exit
					}
					switch choice {
					case 1:

						for {
							fmt.Printf("your user name : %s your password : %s \n Enter new username and password \n", c.Name, c.Password)
							scanner.Scan()
							newUsername := strings.TrimSpace(scanner.Text())
							if newUsername == "0" {
								break
							}
							scanner.Scan()
							newPassword := strings.TrimSpace(scanner.Text())
							if newPassword == "0" {
								break
							}
							for range customers {

								if _, err := os.Stat(c.Name); err == nil {

									err := os.Rename(c.Name, newUsername)
									if err != nil {
										fmt.Println("err while renamign file", err)
									}
								}

								c.Name = newUsername
								c.Password = newPassword
								newUsername = FillDot(c.Name, 10)
								newPassword = FillDot(c.Password, 10)

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
							input := GetInput("Your phone number is " + c.PhoneNumber + ".\nEnter new phone number:")
							if len(input) == 11 {
								c.PhoneNumber = input
								file, err := os.OpenFile("src/customer.txt", os.O_WRONLY, 0644)
								if err != nil {
									fmt.Println("err while opening file", err)
								}
								for _, x := range customers {
									if checkName(input, x.PhoneNumber) == false {
										fmt.Println("This phonenumber already exist")
										break changeLoop
									}
								}
								file.Seek(int64(i*41), io.SeekStart)
								file.WriteString(input)
								fmt.Println("your new phonenumber :", c.PhoneNumber)

								closer([]*os.File{file})
								break
							} else {
								fmt.Println("phonenumber most be 11c")
							}

						}
					case 3:
						fileSelf, err := os.OpenFile("src/"+customers[i].Name, os.O_RDONLY, 0644)
						if err != nil {
							fmt.Println("err while opening file")
							break
						}
						buf := make([]byte, 3)
						var myBook []string
						Line := 0
						for {
							fileSelf.Seek(int64(Line*4), io.SeekStart)
							_, err = fileSelf.Read(buf)
							if err == io.EOF {
								fmt.Println("End")
								closer([]*os.File{fileSelf})
								break
							}
							id, err := strconv.Atoi(strings.Trim(string(buf), "."))
							if err != nil {
								Line++
								continue
							}

							for i, _ := range books {
								if id == books[i].ID {
									myBook = append(myBook, books[i].Name)
									fmt.Println("Found Book: ", books[i].Name)
								}
							}
							Line++
						}
						for i, _ := range myBook {
							fmt.Println(myBook[i])
						}
						closer([]*os.File{fileSelf})
					case 4:
					Loop:
						for {
							if c.Debt > 1000000 {
								fmt.Println("you should pay your debt first ")
								break
							}
							for _, b := range books {
								fmt.Println(b.Name, b.Price, b.ID)
							}
							fileSelf, err := os.OpenFile("src/"+c.Name, os.O_WRONLY|os.O_CREATE, 0644)
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
								input := GetIntInput("Enter IDs [0 for back]: ")
								if input == 0 {
									break Loop
								}
								found := false
								for indexB, b := range books {
									if input != b.ID {
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
										books[indexB].Quantity = books[indexB].Quantity - 1
										customers[i].Debt = customers[i].Debt + books[indexB].Price
										fileSelf.WriteString(FillDot(strconv.Itoa(input), 3) + "\n")
										fileBook.WriteString(FillDot(strconv.Itoa(books[indexB].Quantity), 5))
										fileCustomer.WriteString(FillDot(strconv.Itoa(customers[i].Debt), 8))
										closer([]*os.File{fileSelf, fileBook, fileCustomer})
									}
								}
								if found != true {
									fmt.Println("unknown id")
								}
							}

						}
					case 5:
						for {
							fmt.Println("you can pay :", customers[i].Debt)
							input := GetMenu("Debt", []string{"pay"})
							if input == 0 {
								break
							}
							if input == 1 {
								customers[i].Debt = 0
								fileCustomer, err := os.OpenFile("src/customer.txt", os.O_WRONLY, 0644)
								if err != nil {
									fmt.Println("err while opening file")
									break
								}
								fileCustomer.Seek(int64(i*41)+31, io.SeekStart)
								fileCustomer.WriteString("0.......")
								closer([]*os.File{fileCustomer})
							}

						}
					}

				}
			}
		}
	}
}
