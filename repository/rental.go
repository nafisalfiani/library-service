package repository

import (
	"fmt"
	"library/entity"

	"gorm.io/gorm"
)

type rental struct {
	db *gorm.DB
}

//go:generate mockery --name Rental
type RentalInterface interface {
	ListOutstanding() ([]entity.Rental, error)
	ListHistory() ([]entity.Rental, error)
	Create(rental entity.Rental, paymentMethod string) (entity.Rental, error)
	Delete(rentalId int) error
}

// initRental create rental repository
func initRental(db *gorm.DB) RentalInterface {
	return &rental{
		db: db,
	}
}

func (r *rental) ListOutstanding() ([]entity.Rental, error) {
	rentals := []entity.Rental{}
	if err := r.db.Find(&rentals, r.db.Where("status = 'active'")).Error; err != nil {
		return rentals, errorAlias(err)
	}

	return rentals, nil
}

func (r *rental) ListHistory() ([]entity.Rental, error) {
	rentals := []entity.Rental{}
	if err := r.db.Find(&rentals, r.db.Where("status = 'closed'")).Error; err != nil {
		return rentals, errorAlias(err)
	}

	return rentals, nil
}

func (r *rental) Create(rental entity.Rental, paymentMethod string) (entity.Rental, error) {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		// creating payment for its Id
		payment := entity.Payment{
			UserId: rental.UserId,
		}
		if err := tx.Create(&payment).Error; err != nil {
			return err
		}

		for i := range rental.RentalDetails {
			// fetch book info
			book := entity.Book{
				Id: rental.RentalDetails[i].BookId,
			}
			if err := tx.First(&book).Error; err != nil {
				return err
			}

			// creating rental detail
			rental.RentalDetails[i].RentalId = rental.Id
			rental.RentalDetails[i].BookName = book.Name
			rental.RentalDetails[i].RentalCost = book.RentalCost

			// reduce book stock
			book.StockAvailability--
			if book.StockAvailability < 0 {
				return fmt.Errorf(fmt.Sprintf("%v out of stock", book.Name))
			}
			if err := tx.Save(book).Error; err != nil {
				return err
			}

			// totaling cost
			payment.Amount = payment.Amount + (book.RentalCost * float64(rental.RentalDetails[i].RentalDuration))
		}

		// creating rental
		rental.PaymentId = payment.Id
		if err := tx.Create(&rental).Error; err != nil {
			return err
		}

		// updating payment data
		if err := tx.Save(payment).Error; err != nil {
			return err
		}

		rental.Payment = &payment

		return nil
	}); err != nil {
		return rental, errorAlias(err)
	}

	return rental, nil
}

func (r *rental) Delete(rentalId int) error {
	if err := r.db.Delete(entity.Rental{Id: rentalId}).Error; err != nil {
		return errorAlias(err)
	}

	return nil
}
