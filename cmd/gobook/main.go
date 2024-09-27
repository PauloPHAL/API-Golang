package main

import (
	"database/sql"
	"gobooks/internal/cli"
	"gobooks/internal/services"
	"gobooks/internal/web"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "admin:admin@tcp(34.72.220.48:3306)/books")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	bookService := services.NewBookService(db)
	bookController := web.NewBookController(bookService)

	if len(os.Args) > 1 && (os.Args[1] == "search" || os.Args[1] == "simulate") {
		bookCLI := cli.NewBookCli(bookService)
		bookCLI.Run()
		return
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /books", bookController.GetAllBooks)
	router.HandleFunc("POST /books", bookController.CreateBook)
	router.HandleFunc("GET /books/{id}", bookController.GetBookByID)
	router.HandleFunc("PUT /books/{id}", bookController.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", bookController.DeleteBook)

	//desafio fazer a leitura e o simulate via http

	//router.Handle("GET /books/search/{title}", bookController.SearchBook)
	//router.Handle("POST /books/simulate", bookController.SimulateBook)

	http.ListenAndServe(":8080", router)
}
