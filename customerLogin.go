package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func editFile(fieldIndex int, lineIndex int, newStr string, file *os.File) (int, error) {
	var index int64
	var typee int
	if lineIndex-1 < 0 {
		index = 0
	} else {
		index = customers[lineIndex-1].Offset
	}
	file.Seek(index, io.SeekStart)
	line, err := readLineFromFile(file)
	if err != nil {
		return 0, err
	}
	fields := strings.Split(line, "|")
	switch fieldIndex {
	case 1:
		if len(newStr) != 11 {
			return 1, fmt.Errorf("Wrong number")
		}
		fields[0] = newStr
		typee = 1
	case 2:
		fields[1] = newStr
		typee = 2
	case 3:
		fields[2] = newStr
		typee = 3
	case 4:
		fields[3] = newStr
		typee = 4
	}
	file.Seek(index, io.SeekStart)
	file.WriteString(fields[0] + "|" + fields[1] + "|" + fields[2] + "|" + fields[3])

	return typee, nil
}
func customerLogin() {
	for {
		phoneNumber := GetInput("Enter phone number :")
		password := GetInput("Enter password")
		for i, c := range customers {
			if customers[i].PhoneNumber == phoneNumber && c.Password == password {
				fmt.Println("Welcome" + c.Name)
				for {
					input := GetMenu("Customer panel", []string{"Edit profile", "Pay debt", "New book", "My books"})
					switch input {
					case 1:
						input := GetMenu("Edit profile", []string{"PhoneNumber", "Name", "Password"})
						for {
							newStr := GetInput("Enter new string")

							file, err := os.OpenFile("src/customer.txt", os.O_RDWR, 0644)
							if err != nil {
								fmt.Println(err)
								break
							}
							typee, errrr := editFile(input, i, newStr, file)
							switch typee {
							case 1:
								c.PhoneNumber = newStr
							case 2:
								c.Name = newStr
							case 3:
								c.Password = newStr
							}
							if err != nil {
								fmt.Println(errrr)
							}
							break

						}

					case 2:

					case 3:

					case 4:

					}
				}
			}
		}
	}
}
