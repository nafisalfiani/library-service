package repository

import (
	"library/entity"

	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

type UserInterface interface {
	List() ([]entity.User, error)
	Get(username string) (entity.User, error)
	Create(user entity.User) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	Delete(userId int) error
}

// initUser create user repository
func initUser(db *gorm.DB) UserInterface {
	return &user{
		db: db,
	}
}

// List returns list of users
func (s *user) List() ([]entity.User, error) {
	users := []entity.User{}
	if err := s.db.Find(&users).Error; err != nil {
		return users, errorAlias(err)
	}

	return users, nil
}

// Get returns specific user by username
func (s *user) Get(username string) (entity.User, error) {
	user := entity.User{}
	if err := s.db.First(&user, &entity.User{Username: username}).Error; err != nil {
		return user, errorAlias(err)
	}

	return user, nil
}

// Create creates new data
func (s *user) Create(user entity.User) (entity.User, error) {
	if err := s.db.Create(&user).Error; err != nil {
		return user, errorAlias(err)
	}

	return user, nil
}

// Update updates existing data
func (s *user) Update(user entity.User) (entity.User, error) {
	if err := s.db.Save(&user).Error; err != nil {
		return user, errorAlias(err)
	}

	return user, nil
}

// Delete deletes existing data
func (s *user) Delete(userId int) error {
	if err := s.db.Delete(&entity.User{Id: userId}).Error; err != nil {
		return errorAlias(err)
	}

	return nil
}
