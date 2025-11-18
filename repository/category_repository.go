package repository

import "petshop/model"

type CategoryRepository interface {
	FindAll() ([]model.Category, error)
	FindById(categoryId int) (category model.Category, err error)
	Save(category model.Category) error
	Update(category model.Category) error
	Delete(categoryId int) error
}