package main

import (
	"fmt"
	"io"
	"os"
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
			if c.phonenumber == phonenumber {

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

							file.Close()
							break
						}
					}
				case 2:

				case 3:

				case 4:
				}
			}
		}
	}
}
