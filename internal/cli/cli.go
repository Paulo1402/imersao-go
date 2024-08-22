package cli

import (
	"fmt"
	"gobooks/internal/service"
	"os"
	"strconv"
	"time"
)

type BookCLI struct {
	BookService *service.BookService
}

func NewBookCLI(bookService *service.BookService) *BookCLI {
	return &BookCLI{BookService: bookService}
}

func (cli *BookCLI) Run() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gobooks <command> [<args>]")
		return
	}

	command := os.Args[1]

	switch command {
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: gobooks search <book_title>")
			return
		}

		bookName := os.Args[2]
		cli.searchBooks(bookName)
	case "simulate":
		if len(os.Args) < 3 {
			fmt.Println("Usage: gobooks simulate <book_id> <book_id> <book_id> ...")
			return
		}

		bookIDs := os.Args[2:]
		cli.simulateReading(bookIDs)
	}
}

func (cli *BookCLI) searchBooks(name string) {
	books, err := cli.BookService.SearchBooksByName(name)

	if err != nil {
		fmt.Println("Error searching books:", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("No books found")
		return
	}

	fmt.Printf("%d book(s) found:\n", len(books))

	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Genre: %s\n", book.ID, book.Title, book.Author, book.Genre)
	}
}

func (cli *BookCLI) simulateReading(bookIDsStr []string) {
	var bookIDs []int

	for _, idStr := range bookIDsStr {
		id, err := strconv.Atoi(idStr)

		if err != nil {
			fmt.Println("Invalid book ID:", idStr)
			continue
		}

		bookIDs = append(bookIDs, id)
	}

	responses := cli.BookService.SimulateMultipleReadings(bookIDs, 2*time.Second)

	for _, response := range responses {
		fmt.Println(response)
	}
}
