package repository

import "petshop/model"

type BranchRepository interface {
	FindAll() ([]model.Branch, error)
	FindById(branchId int) (branch model.Branch, err error)
	Save(branch model.Branch) error
	Update(branch model.Branch) error
	Delete(branchId int) error
}