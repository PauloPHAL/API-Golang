package services

import (
	"database/sql"
	"fmt"
	"time"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Genre  string
}

type BookService struct {
	db *sql.DB
}

func NewBookService(db *sql.DB) *BookService {
	return &BookService{db}
}

func (s *BookService) CreateBook(book *Book) error {
	query := "INSERT INTO books (title, author, genre) VALUES (?,?,?)"
	_, err := s.db.Exec(query, book.Title, book.Author, book.Genre)
	if err != nil {
		return err
	}
	return err
}

func (s *BookService) GetBookById(id int) (*Book, error) {
	query := "SELECT * FROM books WHERE id = ?"
	row := s.db.QueryRow(query, id)
	var book Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (s *BookService) GetAllBooks() ([]Book, error) {
	query := "SELECT * FROM books"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (s *BookService) UpdateBook(book *Book) error {
	query := "UPDATE books SET title = ?, author = ?, genre = ? WHERE id = ?"
	_, err := s.db.Exec(query, book.Title, book.Author, book.Genre, book.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookService) DeleteBook(id int) error {
	query := "DELETE FROM books WHERE id = ?"
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookService) SimulateReading(bookID int, duration time.Duration, results chan<- string) {
	book, err := s.GetBookById(bookID)
	if err != nil || book == nil {
		results <- fmt.Sprintf("Book with ID %d not found", bookID)
		return
	}
	time.Sleep(duration)
	results <- fmt.Sprintf("Finished reading %s by %s", book.Title, book.Author)
}

func (s *BookService) SimulateMultipleReadings(bookIDs []int, duration time.Duration) []string {
	results := make(chan string, len(bookIDs))
	for _, id := range bookIDs {
		go func(bookID int) {
			s.SimulateReading(bookID, duration, results)
		}(id)
	}

	var responses []string
	// for res := range results {
	// 	responses = append(responses, res)
	// }

	for range bookIDs {
		responses = append(responses, <-results) // pause until a value is received
	}

	close(results)
	return responses
}

func (s *BookService) SearchBooksByName(name string) ([]Book, error) {
	sql := "SELECT * FROM books WHERE title like ?"

	rows, err := s.db.Query(sql, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Author, &book.Title, &book.Genre)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
