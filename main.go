package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var books []Book
var customers []Customer
var staffs []Staff
var scanner = bufio.NewScanner(os.Stdin)

func GetMenu(title string, items []string) int {
	fmt.Printf("[%s]\n", title)
	for i, item := range items {
		fmt.Printf("  %d. %s\n", i+1, item)
	}
	fmt.Println("----------------------")
	for {
		fmt.Println("Enter menu index: ")
		scanner.Scan()
		indexStr := scanner.Text()
		if indexStr == "" {
			return 0
		}
		if index, err := strconv.Atoi(strings.TrimSpace(indexStr)); err != nil {
			fmt.Printf("Please enter index between %d and %d\n", 1, len(items))
		} else {
			if index < 0 || index > len(items) {
				fmt.Printf("Please enter index between %d and %d\n", 1, len(items))
			} else {
				return index
			}
		}
	}
}

func GetInput(title string) string {
	fmt.Println(title)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func GetIntInput(title string) int {
	for {
		n, err := strconv.Atoi(GetInput(title))
		if err == nil {
			return n
		}
	}
}

func FillDot(input string, l int) string {
	delta := l - len(input)
	return input + strings.Repeat(".", delta)
}
func closer(files []*os.File) {
	for _, f := range files {
		f.Close()
	}
}

func checkName(name string, otherNames string) bool {
	if name == otherNames {
		return false
	}
	return true
}

func main() {
	readBook()
	readCustomer()
	readStaff()

	fmt.Println("You can back every time by entering 0")

	for {
		choice := GetMenu("Login", []string{"Customer", "Staff", "Exit"})
		switch choice {
		case 1:
			customerLogin()
		case 2:
			staffLogin()
		case 0, 3:
			return
		}
	}
}
