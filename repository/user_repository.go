package repository

import "petshop/model"

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindById(userId int) (user model.User, err error)
	Save(user model.User) error
	Update(user model.User) error
	Delete(userId int) error
	FindByEmail(email string) (*model.User, error)
}