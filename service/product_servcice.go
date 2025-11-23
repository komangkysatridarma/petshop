package service

import (
	"petshop/data/request"
	"petshop/data/response"
)

type ProductService interface {
	FindAll() ([]response.ProductResponse, error)
	FindById(productId int) (response.ProductResponse, error)
	Create(request.CreateProductRequest) error
	Update(request.UpdateProductRequest) error
	Delete(productId int) error
}
