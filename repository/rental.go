package repository

import (
	"library/entity"

	"gorm.io/gorm"
)

type rental struct {
	db *gorm.DB
}

type RentalInterface interface {
	List() ([]entity.Rental, error)
	Create(rental entity.Rental) (entity.Rental, error)
	Delete(rentalId int) error
}

// initRental create rental repository
func initRental(db *gorm.DB) RentalInterface {
	return &rental{
		db: db,
	}
}

func (r *rental) List() ([]entity.Rental, error) {
	return nil, nil
}

func (r *rental) Create(rental entity.Rental) (entity.Rental, error) {
	return rental, nil
}

func (r *rental) Delete(rentalId int) error {
	return nil
}
