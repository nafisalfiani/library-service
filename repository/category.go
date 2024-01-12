package repository

import (
	"library/entity"

	"gorm.io/gorm"
)

type category struct {
	db *gorm.DB
}

//go:generate mockery --name Category
type CategoryInterface interface {
	Get(category entity.Category) (entity.Category, error)
	Create(category entity.Category) (entity.Category, error)
}

// initCategory create category repository
func initCategory(db *gorm.DB) CategoryInterface {
	return &category{
		db: db,
	}
}

func (c *category) Get(category entity.Category) (entity.Category, error) {
	if err := c.db.First(&category, c.db.Where("name = ?", category.Name)).Error; err != nil {
		return category, errorAlias(err)
	}

	return category, nil
}

func (c *category) Create(category entity.Category) (entity.Category, error) {
	if err := c.db.Create(&category).Error; err != nil {
		return category, errorAlias(err)
	}

	return category, nil
}
