package service

import (
	"petshop/data/request"
	"petshop/data/response"
)

type BranchService interface {
	FindAll() (branches []response.BranchResponse, err error)
	FindById(branchId int) (branch response.BranchResponse, err error)
	Create(branch request.CreateBranchRequest) error
	Update(branch request.UpdateBranchRequest) error
	Delete(branchId int) error
}
