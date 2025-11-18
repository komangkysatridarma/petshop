package repository

import (
	"errors"
	"petshop/data/request"
	"petshop/model"

	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	Db *gorm.DB
}

func NewRoleRepositoryImpl(Db *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{Db: Db}
}

func (r RoleRepositoryImpl) FindAll() (roles []model.Role, err error) {
	results := r.Db.Find(&roles)
	if results.Error != nil {
		return nil, results.Error
	}
	return roles, nil
}

func (r *RoleRepositoryImpl) FindById(roleId int) (role model.Role, err error) {
	result := r.Db.First(&role, roleId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.Role{}, errors.New("role not found")
		}
		return model.Role{}, result.Error
	}
	return role, nil
}

func (r *RoleRepositoryImpl) Save(role model.Role) error {
	result := r.Db.Create(&role)
	return result.Error
}

func (r *RoleRepositoryImpl) Update(role model.Role) error {
	var data = request.UpdateRoleRequest{
		Name: role.Name,
	}
	result := r.Db.Model(&role).Updates(data)
	return result.Error
}

func (r *RoleRepositoryImpl) Delete(roleId int) error {
	var role model.Role
	result := r.Db.Where("id = ?", roleId).Delete(&role)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("role not found")
	}
	return nil
}