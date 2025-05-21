package repository

import (
	"sync"

	"katalis.ai-home-test/entity"
)

var (
	books   = make(map[string]entity.Book)
	booksDB sync.RWMutex
)

func GetAllBooks() []entity.Book {
	booksDB.RLock()
	defer booksDB.RUnlock()
	all := make([]entity.Book, 0, len(books))
	for _, b := range books {
		all = append(all, b)
	}
	return all
}

func GetBook(isbn string) (entity.Book, bool) {
	booksDB.RLock()
	defer booksDB.RUnlock()
	b, ok := books[isbn]
	return b, ok
}

func CreateBook(b entity.Book) bool {
	booksDB.Lock()
	defer booksDB.Unlock()
	if _, exists := books[b.ISBN]; exists {
		return false
	}
	books[b.ISBN] = b
	return true
}

func UpdateBook(isbn string, b entity.Book) bool {
	booksDB.Lock()
	defer booksDB.Unlock()
	if _, exists := books[isbn]; !exists {
		return false
	}
	b.ISBN = isbn
	books[isbn] = b
	return true
}

func DeleteBook(isbn string) bool {
	booksDB.Lock()
	defer booksDB.Unlock()
	if _, exists := books[isbn]; !exists {
		return false
	}
	delete(books, isbn)
	return true
}
