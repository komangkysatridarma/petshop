package repository

import (
	"errors"
	"petshop/data/request"
	"petshop/model"

	"gorm.io/gorm"
)

type BranchRepositoryImpl struct {
	Db *gorm.DB
}

func NewBranchRepositoryImpl(Db *gorm.DB) BranchRepository {
	return &BranchRepositoryImpl{Db: Db}
}

func (b BranchRepositoryImpl) FindAll() (branches []model.Branch, err error) {
	results := b.Db.Find(&branches)
	if results.Error != nil {
		return nil, results.Error
	}
	return branches, nil
}

func (b *BranchRepositoryImpl) FindById(branchId int) (branch model.Branch, err error) {
	result := b.Db.Find(&branch, branchId)
	if result.Error != nil {
		return model.Branch{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.Branch{}, errors.New("branch is not found")
	}

	return branch, nil
}

func (b *BranchRepositoryImpl) Save(branch model.Branch) error {
	result := b.Db.Create(&branch)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *BranchRepositoryImpl) Update(branch model.Branch) error {
	var data = request.UpdateBranchRequest{
		Name:     branch.Name,
		Code:     branch.Code,
		Address:  branch.Address,
		Phone:    branch.Phone,
		Timezone: branch.Timezone,
	}

	result := b.Db.Model(&branch).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *BranchRepositoryImpl) Delete(branchId int) error {
	var branch model.Branch
	result := b.Db.Where("id = ?", branchId).Delete(&branch)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
