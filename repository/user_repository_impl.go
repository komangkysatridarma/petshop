package repository

import (
	"errors"
	"petshop/data/request"
	"petshop/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

func (u *UserRepositoryImpl) FindAll() (users []model.User, err error) {
	results := u.Db.Find(&users)
	if results.Error != nil {
		return nil, results.Error
	}
	return users, nil
}

func (u *UserRepositoryImpl) FindById(userId int) (user model.User, err error) {
	result := u.Db.Find(&user, userId)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.User{}, errors.New("user is not found")
	}
	return user, nil
}

func (u *UserRepositoryImpl) Save(user model.User) error {
	result := u.Db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepositoryImpl) Update(user model.User) error {
	var data = request.UpdateUserRequest{
		Id:           user.Id,
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		Role:         user.Role,
		Phone_number: user.Phone_number,
		Branch_id:    user.Branch_id,
	}

	result := u.Db.Model(&user).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepositoryImpl) Delete(userId int) error {
	var user model.User

	result := u.Db.Where("id = ?", userId).Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepositoryImpl) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.Db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
