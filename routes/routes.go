package routes

import (
	"net/http"

	"katalis.ai-home-test/handler"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/books", handler.BooksHandler)
	mux.HandleFunc("/books/", handler.BookByISBNHandler)
}
