package repository

import "petshop/model"

type ProductRepository interface {
	FindAll() ([]model.Product, error)
	FindById(productId int) (product model.Product, err error)
	FindByCode(code string) (product model.Product, err error)
	Save(product model.Product) error
	Update(product model.Product) error
	Delete(productId int) error
}
