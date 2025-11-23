package repository

import (
	"errors"
	"petshop/model"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductRepositoryImpl(Db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{Db: Db}
}

func (p *ProductRepositoryImpl) FindAll() (products []model.Product, err error) {
	results := p.Db.Find(&products)
	if results.Error != nil {
		return nil, results.Error
	}
	return products, nil
}

func (p *ProductRepositoryImpl) FindById(productId int) (product model.Product, err error) {
	results := p.Db.First(&product, productId)
	if results.Error != nil {
		return model.Product{}, results.Error
	}

	if results.RowsAffected == 0 {
		return model.Product{}, errors.New("product is not found")
	}

	return product, nil
}

func (p *ProductRepositoryImpl) FindByCode(code string) (product model.Product, err error) {
	result := p.Db.Where("code = ?", code).First(&product)
	if result.Error != nil {
		return model.Product{}, result.Error
	}
	return product, nil
}

func (p *ProductRepositoryImpl) Save(product model.Product) error {
	results := p.Db.Create(&product)
	if results.Error != nil {
		return results.Error
	}
	return nil
}

func (p *ProductRepositoryImpl) Update(product model.Product) error {
	results := p.Db.Save(&product)
	if results.Error != nil {
		return results.Error
	}
	return nil
}

func (p *ProductRepositoryImpl) Delete(productId int) error {
	var product model.Product
	results := p.Db.Where("id = ?", productId).Delete(&product)
	if results.Error != nil {
		return results.Error
	}
	return nil
}
