package service

import (
	"petshop/repository"

	"github.com/go-playground/validator/v10"
)

type BranchServiceImpl struct {
	BranchRepository repository.BranchRepository
	Validate         *validator.Validate
}
