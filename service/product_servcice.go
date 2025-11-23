package service

import (
	"petshop/data/request"
	"petshop/data/response"
)

type ProductService interface {
	FindAll() ([]response.ProductResponse, error)
	FindById(productId int) (response.ProductResponse, error)
	Create(product request.CreateProductRequest) error
	Update(product request.UpdateProductRequest) error
	Delete(productId int) error
}
