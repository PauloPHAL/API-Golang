package web

import (
	"encoding/json"
	"gobooks/internal/services"
	"net/http"
	"strconv"
)

type BookController struct {
	bookService *services.BookService
}

func NewBookController(bookService *services.BookService) *BookController {
	return &BookController{bookService}
}

func (c *BookController) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := c.bookService.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book services.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.bookService.CreateBook(&book)
	if err != nil {
		http.Error(w, "falided to create book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid book ID", http.StatusBadRequest)
		return
	}

	var book services.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	book.ID = id

	if err := h.bookService.UpdateBook(&book); err != nil {
		http.Error(w, "failed to update book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (h *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid book ID", http.StatusBadRequest)
		return
	}

	if err := h.bookService.DeleteBook(id); err != nil {
		http.Error(w, "failed to delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *BookController) GetBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := h.bookService.GetBookById(id)
	if err != nil {
		http.Error(w, "failed to get book", http.StatusInternalServerError)
		return
	}
	if book == nil {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// func (h *BookController) SearchBook(w http.ResponseWriter, r *http.Request) {
// 	title := r.PathValue("title")
// 	books, err := h.bookService.SearchBook(title)
// 	if err != nil {
// 		http.Error(w, "failed to search book", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(books)
// }
