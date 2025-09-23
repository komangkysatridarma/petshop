package service

import (
	"errors"
	"petshop/data/request"
	"petshop/data/response"
	"petshop/model"
	"petshop/repository"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) (service UserService, err error) {
	if validate == nil {
		return nil, errors.New("validator instance cannot be nil")
	}
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}, err
}

func (u UserServiceImpl) FindAll() (users []response.UserResponse, err error) {
	result, err := u.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}

	for _, value := range result {
		user := response.UserResponse{
			Id:           value.Id,
			Name:         value.Name,
			Email:        value.Email,
			Password:     value.Password,
			Role:         value.Role,
			Phone_number: value.Phone_number,
			Branch_id:    value.Branch_id,
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserServiceImpl) FindById(userId int) (user response.UserResponse, err error) {
	data, err := u.UserRepository.FindById(userId)
	if err != nil {
		return response.UserResponse{}, err
	}

	res := response.UserResponse{
		Id:           data.Id,
		Name:         data.Name,
		Email:        data.Email,
		Password:     data.Password,
		Role:         data.Role,
		Phone_number: data.Phone_number,
		Branch_id:    data.Branch_id,
	}
	return res, nil
}

func (u *UserServiceImpl) Create(user request.CreateUserRequest) (err error) {
	err = u.Validate.Struct(user)

	if err != nil {
		return err
	}

	m := model.User{
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		Role:         user.Role,
		Phone_number: user.Phone_number,
		Branch_id:    user.Branch_id,
	}
	u.UserRepository.Save(m)

	return nil
}

func (u *UserServiceImpl) Update(user request.UpdateUserRequest) (err error) {
	data, err := u.UserRepository.FindById(user.Id)

	if err != nil {
		return err
	}

	data.Name = user.Name
	data.Email = user.Email
	data.Password = user.Password
	data.Role = user.Role
	data.Phone_number = user.Phone_number
	data.Branch_id = user.Branch_id
	u.UserRepository.Update(data)
	return nil
}

func (u *UserServiceImpl) Delete(userId int) (err error) {
	err = u.UserRepository.Delete(userId)

	if err != nil {
		return err
	}
	return nil
}
