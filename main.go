package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var book []Book
var customer []Customer
var staff []Staff

func GetMenu(scanner *bufio.Scanner, title string, items []string) int {
	fmt.Printf("[%s]\n", title)
	for i, item := range items {
		fmt.Printf("  %d. %s\n", i+1, item)
	}
	fmt.Println("----------------------")
	for {
		fmt.Println("Enter menu index: ")
		scanner.Scan()
		indexStr := scanner.Text()
		if index, err := strconv.Atoi(strings.TrimSpace(indexStr)); err != nil {
			fmt.Printf("Please enter index between %d and %d\n", 1, len(items))
		} else {
			if index < 1 || index > len(items) {
				fmt.Printf("Please enter index between %d and %d\n", 1, len(items))
			} else {
				return index
			}
		}
	}
}

func GetInput(scanner *bufio.Scanner, title string) string {
	fmt.Println(title)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func FillDot(input string, l int) string {
	delta := l - len(input)
	return input + strings.Repeat(".", delta)
}

func main() {
	read_book()
	read_customer()
	read_staff()
	scanner := bufio.NewScanner(os.Stdin)
	for {

		c := GetMenu(scanner, "Login", []string{"Customer", "Staff"})

		if c == 1 {
			username := GetInput(scanner, "Enter username:")
			password := GetInput(scanner, "Enter password:")

			found := false
			for i := 0; i < len(customer); i++ {
				if customer[i].name == username && customer[i].password == password {
					fmt.Println("welcome", customer[i].name)
					for {
						found = true

						c := GetMenu(scanner, "Customer Panel", []string{"Change username and password", "My books", "New book", "Pay debt", "Back"})

						if c == 1 {
							for {
								fmt.Println("your username :", customer[i].name, "\n", "your password : ", customer[i].password)
								newUsername := GetInput(scanner, "Enter new username:")
								newPassword := GetInput(scanner, "Enter new password:")

								file, err := os.OpenFile("src/customer.txt", os.O_WRONLY, 0644)
								if err != nil {
									fmt.Println("err while opening customer.txt")
									break
								}
								defer file.Close()
								file.Seek(int64((39*i)+11), io.SeekStart)
								if len(newPassword) > 10 || len(newUsername) > 10 {
									fmt.Println("to long password or username(max = 10c)")
								}
								newUsername = FillDot(newUsername, 10)
								newPassword = FillDot(newPassword, 10)

								for j := 0; j < len(customer); j++ {
									if newPassword == customer[j].name {
										fmt.Println("this username already owned")
										break
									}
								}

								file.WriteString(newUsername)
								file.WriteString(newPassword)
								k := customer[i].name
								customer[i].name = newUsername
								customer[i].password = newPassword
								err = os.Rename("src/"+k, "src/"+customer[i].name)
								if err != nil {
									fmt.Println("err while renaming file :", err)
								}

								fmt.Println("Your username and password successfully changed.")
								break
							}
						} else if c == 2 {
							file, err := os.OpenFile("src/"+customer[i].name, os.O_RDONLY, 0644)
							if err != nil {
								fmt.Println("err while opening file to read ")
								break
							}
							defer file.Close()
							Line := 0

							by := make([]byte, 15)
							for {
								file.Seek(int64(Line*17), io.SeekStart)

								_, err = file.Read(by)
								if err == io.EOF {
									fmt.Println("EOF")
								}
								if err != nil {
									break
								}

								fmt.Println(strings.Trim(string(by), "."))
								Line++
							}
						} else if c == 3 {
							if customer[i].debt < 10000000 {
								for {
									for i := 0; i < len(book); i++ {

										if book[i].quantity > 0 {
											fmt.Println(book[i].name, ":", book[i].price)
										}
									}

									A := GetInput(scanner, "Enter name:")
									found2 := false

									for b := 0; b < len(book); b++ {

										if A == book[b].name {
											found2 = true

											file, err := os.OpenFile("src/"+customer[i].name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
											if err != nil {
												fmt.Println("err while opening file to append", err)
												break
											}
											debt, err := os.OpenFile("src/customer.txt", os.O_WRONLY, 0644)
											if err != nil {
												fmt.Println("err while opening file to set a debt", err)
												break
											}
											bookq, err := os.OpenFile("src/book.txt", os.O_WRONLY, 0644)
											if err != nil {
												fmt.Println("err while opening book.txt to edit", err)
											}

											bj := book[b].name
											bj = FillDot(bj, 15)
											book[b].quantity = book[b].quantity - 1

											j := strconv.Itoa(book[b].quantity)

											j = FillDot(j, 5)
											d := strconv.Itoa(book[b].price + customer[i].debt)
											d = FillDot(d, 8)
											if len(d) > 8 {
												fmt.Println("err to have debt")
												break
											}
											_, err = file.WriteString(bj + "\r\n")
											file.Close()
											if err != nil {
												fmt.Println("err while writing to file", err)
											}

											_, err = bookq.Seek(int64((b*30)+15), io.SeekStart)
											if err != nil {
												fmt.Println("err while writing to book")

											}
											bookq.WriteString(j)

											debt.Seek(int64((i*41)+31), io.SeekStart)

											_, err = debt.WriteString(d)
											if err != nil {
												fmt.Println("err while writing debt")

											}
											bookq.Close()
											debt.Close()
											break
										}
									}
									if found2 != true {
										fmt.Println("unknown book")
									}
								}
							} else {
								fmt.Println("you cant get new book you should pay debt first")

							}
						} else if c == 4 {
							for {
								c := GetMenu(scanner, "Your debt", []string{"Pay", "Exit"})

								if c == 1 {
									file, err := os.OpenFile("src/customer.txt", os.O_WRONLY, 0644)
									if err != nil {
										fmt.Println("err while opening file to edit", err)
										break
									}
									defer file.Close()
									_, err = file.Seek(int64((i*41)+31), io.SeekStart)
									if err != nil {
										fmt.Println("err while seeking")
										break
									}
									_, err = file.WriteString("0.......")
									if err != nil {
										fmt.Println("err while writing to file", err)
									}
									customer[i].debt = 0
								} else if c == 2 {
									break
								} else {
									fmt.Println("unknown command")
								}
							}
						} else if c == 5 {
							break
						} else {
							fmt.Println("unknown command")
						}
					}
				}
			}
			if found != true {
				fmt.Println("unknown username or password")
			}
		} else if c == 2 {
			for {
				found := false
				c := GetInput(scanner, "Enter ID:")
				p := GetInput(scanner, "Enter password:")
				for i := 0; i < len(staff); i++ {
					if c == staff[i].id && p == staff[i].password {
						for {
							fmt.Println("welcome ", staff[i].name, " ", staff[i].lastName)
							c := GetMenu(scanner, "Staff Panel", []string{"Customer", "Books", "Back"})
							if c == 1 {
								c := GetMenu(scanner, "Edit customer", []string{"Check", "Remove", "Add", "Back"})
								_ = c
							} else if c == 2 {
								for {
									c := GetMenu(scanner, "Edit books", []string{"Check", "Remove", "Add", "Edit", "Back"})

									if c == 1 {
										for i := 0; i < len(book); i++ {

											if book[i].quantity > 0 {
												fmt.Println(book[i].name, ":", book[i].price)
											}
										}

									} else if c == 2 {

										var rem []string
										for {
											c := GetInput(scanner, "Enter the names (or empty to back):")
											if c == "" {
												break
											}
											rem = append(rem, c)
										}
										mapD := make(map[string]bool)
										for r := 0; r < len(rem); r++ {
											mapD[rem[r]] = true
										}
										for d := 0; d < len(book); d++ {
											if mapD[book[d].name] {
												file, err := os.OpenFile("src/book.txt", os.O_WRONLY, 0644)
												if err != nil {
													fmt.Println("err while opening book ", err)
												}

												file.Seek(int64(d*30), io.SeekStart)

												for n := d; n < len(book)-1; n++ {
													name := book[n+1].name
													quantity := strconv.Itoa(book[n+1].quantity)
													price := strconv.Itoa(book[n+1].price)
													name = FillDot(name, 15)
													quantity = FillDot(quantity, 5)
													price = FillDot(price, 8)
													file.WriteString(name + quantity + price + "\r\n")
												}
												file.Truncate(int64((len(book) - 1) * 30))
												file.Close()

											}
										}

									} else if c == 3 {
										file, err := os.OpenFile("src/book.txt", os.O_WRONLY, 0644)
										if err != nil {
											fmt.Println("err while openning file", err)
										}
										file.Close()
										_ = GetInput(scanner, "Enter name (or empty to back):")
									} else if c == 4 {

									} else if c == 5 {

									} else {
										fmt.Println("unknown command")
									}
								}
							} else if c == 3 {
								break
							} else {
								fmt.Println("unknown command")
							}

							found = true
							break
						}
					}
				}
				if found != true {

					fmt.Println("unknown user or password")
				}
			}
		} else {
			fmt.Println("unknown command")
		}
	}

}
