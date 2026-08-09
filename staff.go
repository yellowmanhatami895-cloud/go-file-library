package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Staff struct {
	id       string
	password string
	name     string
	lastName string
}

func read_staff() {
	file, err := os.OpenFile("src/staff.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("err while opening staff.txt")
		fmt.Println(err)
	}
	defer file.Close()
	Line := 0
	for {
		file.Seek(int64(Line*35), io.SeekStart)
		by := make([]byte, 33)
		_, err = file.Read(by)
		if err == io.EOF {
			fmt.Println("EOF staff")

		}
		if err != nil {
			break
		}
		p := strings.Trim(string(by[5:13]), ".")
		n := strings.Trim(string(by[13:23]), ".")
		ln := strings.Trim(string(by[23:]), ".")
		staff = append(staff, Staff{
			id:       string(by[:5]),
			password: p,
			name:     n,
			lastName: ln,
		})
		Line++
	}

}
