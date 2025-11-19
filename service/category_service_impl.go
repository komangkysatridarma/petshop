package service

import (
	"errors"
	"petshop/data/request"
	"petshop/data/response"
	"petshop/model"
	"petshop/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
}

func NewCategoryServiceImpl(categoryRepository repository.CategoryRepository, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		Validate:           validate,
	}
}

func (c *CategoryServiceImpl) FindAll() (categories []response.CategoryResponse, err error) {
	result, err := c.CategoryRepository.FindAll()
	if err != nil {
		return nil, err
	}

	for _, value := range result {
		category := response.CategoryResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (c *CategoryServiceImpl) FindById(categoryId int) (category response.CategoryResponse, err error) {
	data, err := c.CategoryRepository.FindById(categoryId)
	if err != nil {
		return response.CategoryResponse{}, err
	}

	res := response.CategoryResponse{
		Id:   data.Id,
		Name: data.Name,
	}
	return res, nil
}

func (c *CategoryServiceImpl) Create(category request.CreateCategoryRequest) (err error) {
	err = c.Validate.Struct(category)
	if err != nil {
		return err
	}

	_, err = c.CategoryRepository.FindByName(category.Name)
	if err == nil {
		return errors.New("category name already exists")
	}

	m := model.Category{
		Name: category.Name,
	}
	return c.CategoryRepository.Save(m)
}

func (c *CategoryServiceImpl) Update(category request.UpdateCategoryRequest) (err error) {
	err = c.Validate.Struct(category)
	if err != nil {
		return err
	}

	data, err := c.CategoryRepository.FindById(category.Id)
	if err != nil {
		return err
	}
	existingCategory, err := c.CategoryRepository.FindByName(category.Name)
	if err == nil && existingCategory.Id != category.Id {
		return errors.New("category name already exists")
	}

	data.Name = category.Name
	return c.CategoryRepository.Update(data)
}

func (c *CategoryServiceImpl) Delete(categoryId int) (err error) {
	_, err = c.CategoryRepository.FindById(categoryId)
	if err != nil {
		return err
	}

	return c.CategoryRepository.Delete(categoryId)
}
