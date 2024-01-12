package repository

import (
	"library/entity"

	"gorm.io/gorm"
)

type payment struct {
	db *gorm.DB
}

//go:generate mockery --name Payment
type PaymentInterface interface {
	List(userId int) ([]entity.Payment, error)
	Get(payment entity.Payment) (entity.Payment, error)
	Create(payment entity.Payment) (entity.Payment, error)
	Update(payment entity.Payment) (entity.Payment, error)
}

// initPayment create payment repository
func initPayment(db *gorm.DB) PaymentInterface {
	return &payment{
		db: db,
	}
}

func (p *payment) List(userId int) ([]entity.Payment, error) {
	payments := []entity.Payment{}
	if err := p.db.Find(&payments, p.db.Where("user_id = ?", userId)).Error; err != nil {
		return payments, errorAlias(err)
	}

	return payments, nil
}

func (p *payment) Get(payment entity.Payment) (entity.Payment, error) {
	if err := p.db.First(&payment).Error; err != nil {
		return payment, errorAlias(err)
	}

	return payment, nil
}

func (p *payment) Create(payment entity.Payment) (entity.Payment, error) {
	if err := p.db.Create(&payment).Error; err != nil {
		return payment, errorAlias(err)
	}

	return payment, nil
}

func (p *payment) Update(payment entity.Payment) (entity.Payment, error) {
	if err := p.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&payment).Error; err != nil {
			return err
		}

		user := entity.User{
			Id: payment.UserId,
		}
		if err := tx.First(&user).Error; err != nil {
			return err
		}

		user.DepositAmount = user.DepositAmount + payment.Amount
		if err := tx.Save(&user).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return payment, errorAlias(err)
	}

	return payment, nil
}
