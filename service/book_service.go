package service

import (
	"katalis.ai-home-test/entity"
	"katalis.ai-home-test/repository"
)

func GetAllBooks() []entity.Book {
	return repository.GetAllBooks()
}

func GetBook(isbn string) (entity.Book, bool) {
	return repository.GetBook(isbn)
}

func CreateBook(b entity.Book) bool {
	return repository.CreateBook(b)
}

func UpdateBook(isbn string, b entity.Book) bool {
	return repository.UpdateBook(isbn, b)
}

func DeleteBook(isbn string) bool {
	return repository.DeleteBook(isbn)
}
