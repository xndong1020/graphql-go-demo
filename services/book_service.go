package services

import (
	"acy.com/graphqlgodemo/dtos"
	"acy.com/graphqlgodemo/models"
	"acy.com/graphqlgodemo/repositories"
)


type IBookService interface {
	CreateBook(bookInput *dtos.BookInput) (*models.Book, error)
	UpdateBook(bookInput *dtos.BookInput, id int) error
	DeleteBook(id int) error
	GetOneBook(id int) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
} 

type BookService struct {
	BookRepository *repositories.BookRepository
}

func NewBookService(bookRepository *repositories.BookRepository) *BookService {
	return &BookService{
		BookRepository: bookRepository,
	}
}

// CreateBook implements IBookService
func (b *BookService) CreateBook(bookInput *dtos.BookInput) (*models.Book, error) {
	return b.BookRepository.CreateBook(bookInput)
}

// DeleteBook implements IBookService
func (b *BookService) DeleteBook(id int) error {
	return b.BookRepository.DeleteBook(id)
}

// GetAllBooks implements IBookService
func (b *BookService) GetAllBooks() ([]*models.Book, error) {
	return b.BookRepository.GetAllBooks()
}

// GetOneBook implements IBookService
func (b *BookService) GetOneBook(id int) (*models.Book, error) {
	return b.BookRepository.GetOneBook(id)
}

// UpdateBook implements IBookService
func (b *BookService) UpdateBook(bookInput *dtos.BookInput, id int) error {
	return b.BookRepository.UpdateBook(bookInput, id)
}
