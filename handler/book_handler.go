package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"katalis.ai-home-test/entity"
	"katalis.ai-home-test/middleware"
	"katalis.ai-home-test/service"
)

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		if page < 1 {
			page = 1
		}
		perPage, _ := strconv.Atoi(r.URL.Query().Get("perPage"))
		if perPage < 1 {
			perPage = 10
		}
		all := service.GetAllBooks()
		start := (page - 1) * perPage
		end := start + perPage
		if start > len(all) {
			start = len(all)
		}
		if end > len(all) {
			end = len(all)
		}
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"data": all[start:end],
			"pagination": map[string]interface{}{
				"page":    page,
				"perPage": perPage,
				"total":   len(all),
				"total_pages": func() int {
					if perPage == 0 {
						return 0
					}
					return (len(all) + perPage - 1) / perPage
				}(),
			},
		}
		json.NewEncoder(w).Encode(response)
	case http.MethodPost:
		// buat book baru
		var b entity.Book
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// validasi
		if b.ISBN == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ISBN is required"))
			return
		}
		if !service.CreateBook(b) {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("ISBN must be unique"))
			return
		}
		go func(msg string) { middleware.LogCh <- msg }("Book created: " + b.ISBN)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(b)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func BookByISBNHandler(w http.ResponseWriter, r *http.Request) {
	isbn := strings.TrimPrefix(r.URL.Path, "/books/")
	if isbn == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		b, ok := service.GetBook(isbn)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(b)
	case http.MethodPut:
		var b entity.Book
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !service.UpdateBook(isbn, b) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(b)
	case http.MethodDelete:
		if !service.DeleteBook(isbn) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		go func(msg string) { middleware.LogCh <- msg }("Book deleted: " + isbn)
		w.WriteHeader(http.StatusOK)
		response := map[string]interface{}{
			"message": "Successfully deleted Book ISBN: " + isbn,
		}
		json.NewEncoder(w).Encode(response)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
