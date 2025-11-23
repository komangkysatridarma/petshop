package service

import (
	"petshop/data/request"
	"petshop/data/response"
	"petshop/model"
	"petshop/repository"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	Validate          *validator.Validate
}

func (p *ProductServiceImpl) FindAll() (products []response.ProductResponse, err error) {
	results, err := p.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}

	for _, value := range results {
		product := response.ProductResponse{
			Id:          value.Id,
			Name:        value.Name,
			Code:        value.Code,
			Price:       value.Price,
			Category_id: value.Category_id,
			Is_service:  value.Is_service,
			Track_batch: value.Track_batch,
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *ProductServiceImpl) FindById(productId int) (product response.ProductResponse, err error) {
	data, err := p.ProductRepository.FindById(productId)
	if err != nil {
		return response.ProductResponse{}, err
	}

	res := response.ProductResponse{
		Id:          data.Id,
		Name:        data.Name,
		Code:        data.Code,
		Price:       data.Price,
		Category_id: data.Category_id,
		Is_service:  data.Is_service,
		Track_batch: data.Track_batch,
	}

	return res, nil
}

func (p *ProductServiceImpl) Create(product request.CreateProductRequest) (err error) {
	err = p.Validate.Struct(product)

	if err != nil {
		return err
	}

	m := model.Product{
		Name:        product.Name,
		Code:        product.Code,
		Price:       product.Price,
		Category_id: product.Category_id,
		Is_service:  product.Is_service,
		Track_batch: product.Track_batch,
	}
	p.ProductRepository.Save(m)

	return nil
}

func (p *ProductServiceImpl) Update(product request.UpdateProductRequest) (err error) {
	data, err := p.ProductRepository.FindById(product.Id)
	if err != nil {
		return err
	}

	data.Name = product.Name
	data.Code = product.Code
	data.Price = product.Price
	data.Category_id = product.Category_id
	data.Is_service = product.Is_service
	data.Track_batch = product.Track_batch
	p.ProductRepository.Update(data)

	return nil
}

func (p *ProductServiceImpl) Delete(productId int) (err error) {
	err = p.ProductRepository.Delete(productId)

	if err != nil {
		return nil
	}
	return nil
}
