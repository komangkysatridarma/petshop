package service

import (
	"errors"
	"petshop/data/request"
	"petshop/data/response"
	"petshop/model"
	"petshop/repository"

	"github.com/go-playground/validator/v10"
)

type BranchServiceImpl struct {
	BranchRepository repository.BranchRepository
	Validate         *validator.Validate
}

func NewBranchServiceImpl(branchRepository repository.BranchRepository, validate *validator.Validate)(service BranchService, err error) {
	if validate == nil {
		return nil, errors.New("validator instance cannot be nil")
	}

	return &BranchServiceImpl{
		BranchRepository: branchRepository,
		Validate: validate,
	}, err
}

func (b BranchServiceImpl) FindAll() (branches []response.BranchResponse, err error) {
	result, err := b.BranchRepository.FindAll()

	if err != nil {
		return nil, err
	}

	for _, value := range result {
		branch := response.BranchResponse{
			Id: value.Id,
			Name: value.Name,
			Code: value.Code,
			Address: value.Address,
			Phone: value.Phone,
			Timezone: value.Timezone,
		}
		branches = append(branches, branch)
	}
	return branches, nil
}

func (b *BranchServiceImpl) FindById(branchId int) (branch response.BranchResponse, err error) {
	data, err := b.BranchRepository.FindById(branchId)
	if err != nil {
		return response.BranchResponse{}, err
	}

	res := response.BranchResponse{
		Id: data.Id,
		Name: data.Name,
		Code: data.Code,
		Address: data.Address,
		Phone: data.Phone,
		Timezone: data.Timezone,
	}
	return res, nil
}

func (b *BranchServiceImpl) Create(branch request.CreateBranchRequest) (err error) {
	err = b.Validate.Struct(branch)

	if err != nil {
		return err
	}

	m := model.Branch{
		Name: branch.Name,
		Code: branch.Code,
		Address: branch.Address,
		Phone: branch.Phone,
		Timezone: branch.Timezone,
	}
	b.BranchRepository.Save(m)

	return nil
}

func (b *BranchServiceImpl) Update(branch request.UpdateBranchRequest) (err error){
	data, err := b.BranchRepository.FindById(branch.Id)
	if err != nil {
		return err
	}

	data.Name = branch.Name
	data.Code = branch.Code
	data.Address = branch.Address
	data.Phone = branch.Phone
	data.Timezone = branch.Timezone
	b.BranchRepository.Update(data)

	return nil
}

func (b *BranchServiceImpl) Delete(branchId int) (err error) {
	err = b.BranchRepository.Delete(branchId)
	if err != nil {
		return err
	}
	return nil
}
