package repository

import (
	"library/entity"

	"gorm.io/gorm"
)

type book struct {
	db *gorm.DB
}

//go:generate mockery --name Book
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
	books := []entity.Book{}
	if err := b.db.Find(&books).Error; err != nil {
		return nil, errorAlias(err)
	}

	return books, nil
}

func (b *book) Get(book entity.Book) (entity.Book, error) {
	if err := b.db.First(&book).Error; err != nil {
		return book, errorAlias(err)
	}

	return book, nil
}

func (b *book) Create(book entity.Book) (entity.Book, error) {
	if err := b.db.Create(&book).Error; err != nil {
		return book, errorAlias(err)
	}

	return book, nil
}

func (b *book) Update(book entity.Book) (entity.Book, error) {
	if err := b.db.Save(&book).Error; err != nil {
		return book, errorAlias(err)
	}

	return book, nil
}

func (b *book) Delete(bookId int) error {
	if err := b.db.Delete(entity.Book{Id: bookId}).Error; err != nil {
		return errorAlias(err)
	}

	return nil
}
