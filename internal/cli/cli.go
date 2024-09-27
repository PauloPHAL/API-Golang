package cli

import (
	"fmt"
	"gobooks/internal/services"
	"os"
	"strconv"
	"time"
)

type BookCli struct {
	services *services.BookService
}

func NewBookCli(service *services.BookService) *BookCli {
	return &BookCli{services: service}
}

func (cli *BookCli) searchBooks(name string) {
	books, err := cli.services.SearchBooksByName(name)
	if err != nil {
		fmt.Println("Erro searching books:", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("No books found.")
		return
	}

	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Genre: %s", book.ID, book.Title, book.Author, book.Genre)
	}
}

func (cli *BookCli) SimulateReadind(bookIDsStr []string) {
	var bookIDs []int
	for _, idString := range bookIDsStr {
		id, err := strconv.Atoi(idString)
		if err != nil {
			fmt.Println("Invalide book ID", idString)
			continue
		}
		bookIDs = append(bookIDs, id)
	}

	responses := cli.services.SimulateMultipleReadings(bookIDs, 2*time.Second)

	for _, response := range responses {
		fmt.Println(response)
	}
}

func (cli *BookCli) Run() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: books <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books search <book title>")
			return
		}
		bookName := os.Args[2]
		cli.searchBooks(bookName)
	case "simulate":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books simulate <book_id> <book_id> <book_id> ...")
			return
		}
		booksIDs := os.Args[2:]
		cli.SimulateReadind(booksIDs)
	}
}
