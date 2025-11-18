package service

import "petshop/data/request"
import "petshop/data/response"

type RoleService interface {
	FindAll() ([]response.RoleResponse, error)
	FindById(roleId int) (response.RoleResponse, error)
	Create(role request.CreateRoleRequest) error
	Update(role request.UpdateRoleRequest) error
	Delete(roleId int) error
}