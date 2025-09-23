package service

import (
	"petshop/data/request"
	"petshop/data/response"
)

type UserService interface {
	Create(user request.CreateUserRequest) error
	Update(user request.UpdateUserRequest) error
	Delete(userId int) error
	FindById(userId int) (task response.UserResponse, err error)
	FindAll() (tasks []response.UserResponse, err error)
}