package repositories

import (
	"acy.com/graphqlgodemo/dtos"
	"acy.com/graphqlgodemo/models"
	"gorm.io/gorm"
)

type IBookRepository interface {
	CreateBook(bookInput *dtos.BookInput) (*models.Book, error)
	UpdateBook(bookInput *dtos.BookInput, id int) error
	DeleteBook(id int) error
	GetOneBook(id int) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
}

type BookRepository struct {
	Db *gorm.DB
}

// CreateBook implements IBookRepository
func (b *BookRepository) CreateBook(bookInput *dtos.BookInput) (*models.Book, error) {
	book := &models.Book{
        Title:     bookInput.Title,
        Author:    bookInput.Author,
        Publisher: bookInput.Publisher,
    }
    err := b.Db.Create(&book).Error

    return book, err
}

// DeleteBook implements IBookRepository
func (b *BookRepository) DeleteBook(id int) error {
	book := &models.Book{}
    err := b.Db.Delete(book, id).Error
    return err
}

// GetAllBooks implements IBookRepository
func (b *BookRepository) GetAllBooks() ([]*models.Book, error) {
	books := []*models.Book{}
    err := b.Db.Find(&books).Error
    return books, err
}

// GetOneBook implements IBookRepository
func (b *BookRepository) GetOneBook(id int) (*models.Book, error) {
	book := &models.Book{}
    err := b.Db.Where("id = ?", id).First(book).Error
    return book, err
}

// UpdateBook implements IBookRepository
func (b *BookRepository) UpdateBook(bookInput *dtos.BookInput, id int) error {
	book := models.Book{
        ID:        id,
        Title:     bookInput.Title,
        Author:    bookInput.Author,
        Publisher: bookInput.Publisher,
    }
    err := b.Db.Model(&book).Where("id = ?", id).Updates(book).Error
    return err
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		Db: db,
	}
}