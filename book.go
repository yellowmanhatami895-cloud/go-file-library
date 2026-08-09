package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	name     string
	quantity int
	price    int
}

func read_book() {
	file, err := os.OpenFile("src/book.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("err while opening book.txt")
		fmt.Println(err)
	}
	defer file.Close()
	Line := 0
	for {
		file.Seek(int64(Line*30), io.SeekStart)
		by := make([]byte, 28)
		_, err = file.Read(by)
		if err == io.EOF {
			fmt.Println("EOF book")

		}
		if err != nil {
			break
		}
		n := strings.Trim(string(by[:15]), ".")
		q, err := strconv.Atoi(strings.Trim(string(by[15:20]), "."))
		if err != nil {
			fmt.Println("err cant conv quantity(book) : ", err)
			break
		}
		p, err := strconv.Atoi(strings.Trim(string(by[20:]), "."))
		if err != nil {
			fmt.Println("err cant conv price(book) : ", err)
		}
		book = append(book, Book{
			name:     n,
			quantity: q,
			price:    p,
		})
		Line++
	}

}
