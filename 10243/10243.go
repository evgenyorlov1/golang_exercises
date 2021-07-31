package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var books []Book

type Book struct {
	book_id int
	title   string
	author  string
	subject string
}

func getBook(id int) string {
	for _, v := range books {
		if v.book_id == id {
			return v.title + " by " + v.author
		}
	}
	return "Book not found"
}

func main() {
	var n, q, val int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Scan(&n)
	scanner.Scan()

	for i := 0; i < n; i++ {
		var book Book
		scanner.Scan()
		id, _ := strconv.Atoi(scanner.Text())
		book.book_id = id

		scanner.Scan()
		book.title = scanner.Text()

		scanner.Scan()
		book.author = scanner.Text()

		scanner.Scan()
		book.subject = scanner.Text()

		books = append(books, book)
		scanner.Scan()
	}

	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		fmt.Scan(&val)
		fmt.Println(getBook(val))
	}
}
