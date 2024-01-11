package repository

import (
	"library/entity"

	"gorm.io/gorm"
)

type deposit struct {
	db *gorm.DB
}

type DepositInterface interface {
	List(userId int) (entity.DepositHistories, error)
	Create(deposit entity.DepositHistory) (entity.DepositHistory, error)
}

// initDeposit create deposit repository
func initDeposit(db *gorm.DB) DepositInterface {
	return &deposit{
		db: db,
	}
}

func (d *deposit) List(userId int) (entity.DepositHistories, error) {
	deposits := []entity.DepositHistory{}
	if err := d.db.Find(&deposits).Error; err != nil {
		return nil, err
	}

	return deposits, nil
}

func (d *deposit) Create(deposit entity.DepositHistory) (entity.DepositHistory, error) {

	return deposit, nil
}
