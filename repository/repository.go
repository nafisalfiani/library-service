package repository

import (
	"library/errors"

	"github.com/xendit/xendit-go/v4"
	"gorm.io/gorm"
)

type Repository struct {
	User     UserInterface
	Book     BookInterface
	Category CategoryInterface
	Deposit  DepositInterface
	Rental   RentalInterface
	Payment  PaymentInterface
	Xendit   XenditInterface
}

func InitRepository(db *gorm.DB, xnd *xendit.APIClient) *Repository {
	return &Repository{
		User:     initUser(db),
		Book:     initBook(db),
		Category: initCategory(db),
		Deposit:  initDeposit(db),
		Rental:   initRental(db),
		Payment:  initPayment(db),
		Xendit:   initXendit(xnd),
	}
}

func errorAlias(err error) error {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return errors.ErrNotFound
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return errors.ErrDuplicatedKey
	default:
		return err
	}
}
