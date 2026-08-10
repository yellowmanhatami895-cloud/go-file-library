package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var book []Book
var customer []Customer
var staff []Staff
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
			return 999
		}
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

func GetInput(title string) string {
	fmt.Println(title)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
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

func staffLogin() {

}
func back(input string, inputInt int) bool {
	if input == "" {
		return true
	}
	if inputInt == 999 {
		return true
	}
	return false
}
func checkName(name string, othernames string) bool {
	if name == othernames {
		return false
	}
	return true
}
func main() {
	read_book()
	read_customer()
	read_staff()
	fmt.Println("you can back every time after presing (empty) + enter")
	for {
		choise := GetMenu("Login", []string{"Customer", "Staff"})
		if back(" ", choise) == true {
			break
		}
		switch choise {
		case 1:
			custoemerLogin()
		case 2:
			staffLogin()

		}
	}
}
