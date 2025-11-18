package service

import "petshop/data/request"
import "petshop/data/response"

type CategoryService interface {
	FindAll() ([]response.CategoryResponse, error)
	FindById(categoryId int) (response.CategoryResponse, error)
	Create(category request.CreateCategoryRequest) error
	Update(category request.UpdateCategoryRequest) error
	Delete(categoryId int) error
}