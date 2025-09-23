package service

import (
	"errors"
	"petshop/repository" // pakai repository yang sudah ada
	"petshop/utils"
	"strconv"
)

type AuthService interface {
	Login(email, password string) (string, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{userRepository: userRepository}
}

func (a *authService) Login(email, password string) (string, error) {
	user, err := a.userRepository.FindByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	// cek password (sementara plain text, idealnya bcrypt)
	if user.Password != password {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateToken(strconv.Itoa(user.Id), user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
