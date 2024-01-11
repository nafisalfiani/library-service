package repository

import (
	"library/errors"

	"github.com/xendit/xendit-go/v4"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

type Repository struct {
	User     UserInterface
	Book     BookInterface
	Category CategoryInterface
	Rental   RentalInterface
	Payment  PaymentInterface
	Xendit   XenditInterface
	Mail     MailInterface
}

func InitRepository(db *gorm.DB, xnd *xendit.APIClient, dialer *gomail.Dialer) *Repository {
	return &Repository{
		User:     initUser(db),
		Book:     initBook(db),
		Category: initCategory(db),
		Rental:   initRental(db),
		Payment:  initPayment(db),
		Xendit:   initXendit(xnd),
		Mail:     initMail(dialer),
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
