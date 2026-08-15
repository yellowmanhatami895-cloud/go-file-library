package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func staffLogin() {
Loop:
	for {
		Id := GetIntInput("ID")
		if Id == 0 {
			break
		}
		password := GetInput("password")
		if password == "0" {

		}
		for i, s := range staffs {
			if Id == s.ID && password == staffs[i].Password {
				for {
					fmt.Println("welcome", s.Name+" "+s.LastName)
					input := GetMenu("Staff panel", []string{"Books", "Customers", "Edit self"})
					if input == 0 {
						break Loop
					}
					switch input {
					case 1:
						input = GetMenu("Books", []string{"Add", "Remove", "Edit", "Retrieving the book"})
						if input == 0 {
							break Loop
						}
						switch input {
						case 1:
						addLoop:
							for {
								fileBook, err := os.OpenFile("src/book.txt", os.O_WRONLY, 0644)
								if err != nil {
									fmt.Println("err while opening book")
									break
								}

								found := false

								newBook_name := GetInput("Enter name")
								if newBook_name == "0" {
									closer([]*os.File{fileBook})
									break
								}

								newBook_quantity := GetInput("Enter quantity")
								if newBook_quantity == "0" {
									closer([]*os.File{fileBook})
									break
								}

								newBook_price := GetIntInput("Enter price")
								if newBook_price == 0 {
									closer([]*os.File{fileBook})
									break
								}

								quantity_int, err := strconv.Atoi(newBook_quantity)
								if err != nil {
									fmt.Println("err while conv", err)
									closer([]*os.File{fileBook})
									break
								}

								fillName := FillDot(newBook_name, 15)
								fillQuantity := FillDot(newBook_quantity, 5)
								fillPrice := FillDot(strconv.Itoa(newBook_price), 8)

								newID := 1
								for {
									used := false

									for _, b := range books {
										if b.ID == newID {
											used = true
											break
										}
									}

									if used == false {
										break
									}

									newID++
								}

								NewIDStr := fmt.Sprintf("%03d", newID)
								for i := range books {
									if books[i].ID == 0 {
										found = true

										fileBook.Seek(int64(i*33), io.SeekStart)
										fileBook.WriteString(NewIDStr + fillName + fillQuantity + fillPrice)

										books[i] = Book{
											Name:     newBook_name,
											Quantity: quantity_int,
											Price:    newBook_price,
											ID:       newID,
										}

										break
									}
								}

								if found == false {
									fmt.Println("no free space")

									books = append(books, Book{
										Name:     newBook_name,
										Quantity: quantity_int,
										Price:    newBook_price,
										ID:       newID,
									})
									fileBook.Seek(0, io.SeekEnd)
									fileBook.WriteString("\r\n" + NewIDStr + fillName + fillQuantity + fillPrice)
								}

								closer([]*os.File{fileBook})
								break addLoop
							}
						case 2:
							for {
								for _, b := range books {
									fmt.Println(b.Name, b.Quantity, b.Price, b.ID)
								}
								input := GetIntInput("Enter the ID [0 for back]:")
								if input == 0 {
									break
								}
								for i, b := range books {
									if b.ID == input {

										fileBook, err := os.OpenFile("src/book.txt", os.O_WRONLY, 0644)
										if err != nil {
											fmt.Println("err while opening file", err)
											break
										}
										fileBook.Seek(int64(i*33), io.SeekStart)
										fileBook.WriteString("000...............0....0.......")
										closer([]*os.File{fileBook})
										books[i] = Book{
											Name:     "...............",
											ID:       0,
											Quantity: 0,
											Price:    0,
										}
									}
								}

							}
						case 3:

						case 4:
						}
					case 2:
						input = GetMenu("Customers", []string{"Add", "Remove", "Edit"})
						if input == 0 {
							break Loop
						}
					case 3:
						input = GetMenu("Edit self", []string{"Name", "Last name", "Password"})
						if input == 0 {
							break Loop
						}
					}
				}
			}
		}

	}
}
