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

func main() {
	read_book()
	read_customer()
	read_staff()
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Println("login : customer.1 | staff.2")
		scanner.Scan()
		c := scanner.Text()
		c = strings.TrimSpace(c)

		if c == "1" {

			fmt.Println("enter username")
			scanner.Scan()
			c := scanner.Text()
			c = strings.TrimSpace(c)

			fmt.Println("enter password")
			scanner.Scan()

			p := strings.TrimSpace(scanner.Text())

			found := false
			for i := 0; i < len(customer); i++ {
				if customer[i].name == c && customer[i].password == p {
					fmt.Println("welcome", customer[i].name)
					for {
						found = true
						fmt.Printf("\n change username and password.1 \n My_books.2 \n New book.3 \n pay debt.4 \n Exit.5 \n")
						scanner.Scan()
						c := strings.TrimSpace(scanner.Text())

						if c == "1" {
							for {
								fmt.Println("your username :", customer[i].name, "\n", "your password : ", customer[i].password, "\n", "enter new username and password")
								scanner.Scan()
								c = strings.TrimSpace(scanner.Text())

								scanner.Scan()
								p = strings.TrimSpace(scanner.Text())

								file, err := os.OpenFile("src/customer.txt", os.O_WRONLY, 0644)
								if err != nil {
									fmt.Println("err while opening customer.txt")
									break
								}
								defer file.Close()
								file.Seek(int64((39*i)+11), io.SeekStart)
								if len(p) > 10 || len(c) > 10 {
									fmt.Println("to long password or username(max = 10c)")
								}
								for {
									if len(c) < 10 {
										c = c + "."
									} else {
										break
									}
								}
								for {
									if len(p) < 10 {
										p = p + "."
									} else {
										break
									}
								}

								for j := 0; j < len(customer); j++ {
									if p == customer[j].name {
										fmt.Println("this username already owned")
										break
									}
								}

								file.WriteString(c)
								file.WriteString(p)
								k := customer[i].name
								customer[i].name = c
								customer[i].password = p
								err = os.Rename("src/"+k, "src/"+customer[i].name)
								if err != nil {
									fmt.Println("err while renaming file :", err)
								}

								fmt.Println("your password success fully change \n Exit.1 ")
								fmt.Scan(&c)
								if c == "1" {
									break
								} else {
									fmt.Println("unknown command")
									break
								}

							}
						} else if c == "2" {
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
						} else if c == "3" {
							if customer[i].debt < 10000000 {
								for {
									for i := 0; i < len(book); i++ {

										if book[i].quantity > 0 {
											fmt.Println(book[i].name, ":", book[i].price)
										}
									}

									fmt.Println("enter the name |Exit.2 ")
									scanner.Scan()
									A := scanner.Text()
									A = strings.TrimSpace(A)
									found2 := false
									if A == "2" {
										found2 = true
										break

									}
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
											for {
												if len(bj) < 15 {
													bj = bj + "."
												} else {
													break
												}
											}
											book[b].quantity = book[b].quantity - 1

											j := strconv.Itoa(book[b].quantity)

											for {
												if len(j) < 5 {
													j = j + "."
												} else {
													break
												}
											}
											d := strconv.Itoa(book[b].price + customer[i].debt)

											for {
												if len(d) < 8 {
													d = d + "."
												} else {
													break
												}
											}
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
						} else if c == "4" {
							for {
								fmt.Println("your debt : ", customer[i].debt, "pay.1 | Exit.2")
								scanner.Scan()
								c := scanner.Text()
								c = strings.TrimSpace(c)
								if c == "1" {
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
								} else if c == "2" {
									break
								} else {
									fmt.Println("unknown command")
								}
							}
						} else if c == "5" {
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
		} else if c == "2" {
			for {
				fmt.Println("enter id")
				scanner.Scan()
				c := scanner.Text()
				c = strings.TrimSpace(c)
				found := false
				fmt.Println("enter the  password")
				scanner.Scan()
				p := scanner.Text()
				p = strings.TrimSpace(p)
				for i := 0; i < len(staff); i++ {
					if c == staff[i].id && p == staff[i].password {
						for {
							fmt.Println("welcome ", staff[i].name, " ", staff[i].lastName)
							fmt.Println("customer.1 | books.2 | Exit.3")
							scanner.Scan()
							c := scanner.Text()
							c = strings.TrimSpace(c)
							if c == "1" {
								fmt.Println("check.1 | remove.2 | add.3 |Exit.4")
								scanner.Scan()
								c := scanner.Text()
								c = strings.TrimSpace(c)
							} else if c == "2" {
								for {
									fmt.Println("check.1 | remove.2 | add.3 |Edit.4 |Exit.5")
									scanner.Scan()
									c := scanner.Text()
									c = strings.TrimSpace(c)
									if c == "1" {
										for i := 0; i < len(book); i++ {

											if book[i].quantity > 0 {
												fmt.Println(book[i].name, ":", book[i].price)
											}
										}

									} else if c == "2" {

										var rem []string
										for {
											fmt.Println("Exit.2 |enter the names :")
											scanner.Scan()
											c := scanner.Text()
											c = strings.TrimSpace(c)
											rem = append(rem, c)
											if c == "2" {

												break
											}

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
													quantity := string(book[n+1].quantity)
													price := string(book[n+1].price)
													for {
														if len(name) < 15 {
															name = name + "."
														}
														break
													}
													for {
														if len(quantity) < 5 {
															quantity = quantity + "."
														}
														break
													}
													for {
														if len(price) < 8 {
															price = price + "."
														}
														break
													}
													file.WriteString(name + quantity + price + "\r\n")
												}
												file.Truncate(int64((len(book) - 1) * 30))
												file.Close()

											}
										}

									} else if c == "3" {

									} else if c == "4" {

									} else if c == "5" {

									} else {
										fmt.Println("unknown command")
									}
								}
							} else if c == "3" {
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
