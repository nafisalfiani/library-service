package repository

import (
	"library/entity"

	"gorm.io/gorm"
)

type book struct {
	db *gorm.DB
}

type BookInterface interface {
	List() ([]entity.Book, error)
	Get(book entity.Book) (entity.Book, error)
	Create(book entity.Book) (entity.Book, error)
	Update(book entity.Book) (entity.Book, error)
	Delete(bookId int) error
}

// initBook create book repository
func initBook(db *gorm.DB) BookInterface {
	return &book{
		db: db,
	}
}

func (b *book) List() ([]entity.Book, error) {
	return nil, nil
}

func (b *book) Get(book entity.Book) (entity.Book, error) {
	return book, nil
}

func (b *book) Create(book entity.Book) (entity.Book, error) {
	return book, nil
}

func (b *book) Update(book entity.Book) (entity.Book, error) {
	return book, nil
}

func (b *book) Delete(bookId int) error {
	return nil
}
