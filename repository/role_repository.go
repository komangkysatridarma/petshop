package repository

import "petshop/model"

type RoleRepository interface {
	FindAll() ([]model.Role, error)
	FindById(roleId int) (role model.Role, err error)
	Save(role model.Role) error
	Update(role model.Role) error
	Delete(roleId int) error
}