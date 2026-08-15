package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	ID       int
	Name     string
	Quantity int
	Price    int
}

func readLineFromFile(file *os.File) (string, error) {
	line := make([]byte, 1000)
	readCount, err := file.Read(line)
	if err != nil {
		return "", err
	}
	index := strings.IndexByte(string(line), '\n')
	if index < 0 {
		return string(line[:readCount]), nil
	}
	newSeek := readCount - index - 1
	if newSeek > 0 {
		file.Seek(-int64(newSeek), io.SeekCurrent)
	}

	return string(line[:index-1]), nil
}

func readBookFromFile(file *os.File) (Book, error) {
	line, err := readLineFromFile(file)
	if err != nil {
		return Book{}, err
	}

	fields := strings.Split(line, "|")
	if len(fields) != 4 {
		return Book{}, fmt.Errorf("invalid record line")
	}

	var book Book
	book.ID, err = strconv.Atoi(fields[0])
	if err != nil {
		return Book{}, err
	}

	book.Name = fields[1]

	book.Quantity, err = strconv.Atoi(fields[2])
	if err != nil {
		fmt.Println("Error cant conv quantity(book) : ", err)
		return Book{}, err
	}

	book.Price, err = strconv.Atoi(fields[3])
	if err != nil {
		fmt.Println("err cant conv price(book) : ", err)
		return Book{}, err
	}

	return book, nil
}

func readBook() {
	file, err := os.OpenFile("src/book.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error while opening book.txt", err)
		return
	}
	defer file.Close()

	for {
		book, err := readBookFromFile(file)
		if err != nil {
			break
		}

		books = append(books, book)
	}
}
