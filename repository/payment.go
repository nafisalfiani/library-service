package repository

import (
	"library/entity"

	"gorm.io/gorm"
)

type payment struct {
	db *gorm.DB
}

type PaymentInterface interface {
	List() ([]entity.Payment, error)
	Get(payment entity.Payment) (entity.Payment, error)
	Create(payment entity.Payment) (entity.Payment, error)
}

// initPayment create payment repository
func initPayment(db *gorm.DB) PaymentInterface {
	return &payment{
		db: db,
	}
}

func (p *payment) List() ([]entity.Payment, error) {
	return nil, nil
}

func (p *payment) Get(payment entity.Payment) (entity.Payment, error) {
	return payment, nil
}

func (p *payment) Create(payment entity.Payment) (entity.Payment, error) {
	return payment, nil
}
