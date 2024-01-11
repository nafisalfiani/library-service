package repository

import (
	"library/entity"

	"gorm.io/gorm"
)

type category struct {
	db *gorm.DB
}

type CategoryInterface interface {
	List() ([]entity.Category, error)
	Get(category entity.Category) (entity.Category, error)
	Create(category entity.Category) (entity.Category, error)
}

// initCategory create category repository
func initCategory(db *gorm.DB) CategoryInterface {
	return &category{
		db: db,
	}
}

func (c *category) List() ([]entity.Category, error) {
	return nil, nil
}

func (c *category) Get(category entity.Category) (entity.Category, error) {
	return category, nil
}

func (c *category) Create(category entity.Category) (entity.Category, error) {
	return category, nil
}
