package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Staff struct {
	ID       int
	Password string
	Name     string
	LastName string
}

func readStaffFromFile(file *os.File) (Staff, error) {
	line, err := readLineFromFile(file)
	if err != nil {
		return Staff{}, err
	}
	fields := strings.Split(line, "|")
	var staff Staff
	if len(fields) < 4 {
		return staff, err
	}
	staff.ID, err = strconv.Atoi(fields[0])
	if err != nil {
		return Staff{}, err
	}
	staff.Password = fields[1]
	staff.Name = fields[2]
	staff.LastName = fields[3]
	return staff, nil
}
func readStaff() {
	file, err := os.OpenFile("src/staff.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("err while opening staff.txt")
		fmt.Println(err)
	}
	defer file.Close()

	for {

		staff, err := readStaffFromFile(file)
		if err != nil {
			break
		}
		staffs = append(staffs, staff)
		fmt.Println(staff)
	}

}
