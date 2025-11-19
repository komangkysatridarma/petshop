package repository

import (
    "errors"
    "petshop/model"

    "gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
    Db *gorm.DB
}

func NewCategoryRepositoryImpl(Db *gorm.DB) CategoryRepository {
    return &CategoryRepositoryImpl{Db: Db}
}

func (c *CategoryRepositoryImpl) FindAll() (categories []model.Category, err error) {
    results := c.Db.Find(&categories)
    if results.Error != nil {
        return nil, results.Error
    }
    return categories, nil
}

func (c *CategoryRepositoryImpl) FindById(categoryId int) (category model.Category, err error) {
    result := c.Db.First(&category, categoryId)
    if result.Error != nil {
        return model.Category{}, result.Error
    }

    if result.RowsAffected == 0 {
        return model.Category{}, errors.New("category is not found")
    }

    return category, nil
}

func (c *CategoryRepositoryImpl) FindByName(name string) (category model.Category, err error) {
    result := c.Db.Where("name = ?", name).First(&category)
    if result.Error != nil {
        return model.Category{}, result.Error
    }
    return category, nil
}

func (c *CategoryRepositoryImpl) Save(category model.Category) error {
    result := c.Db.Create(&category)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (c *CategoryRepositoryImpl) Update(category model.Category) error {
    result := c.Db.Save(&category)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (c *CategoryRepositoryImpl) Delete(categoryId int) error {
    var category model.Category
    result := c.Db.Where("id = ?", categoryId).Delete(&category)
    if result.Error != nil {
        return result.Error
    }
    return nil
}