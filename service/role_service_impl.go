package service

import (
	"petshop/data/request"
	"petshop/data/response"
	"petshop/model"
	"petshop/repository"

	"github.com/go-playground/validator/v10"
)

type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
	Validate       *validator.Validate
}

func NewRoleServiceImpl(roleRepository repository.RoleRepository, validate *validator.Validate) RoleService {
	return &RoleServiceImpl{
		RoleRepository: roleRepository,
		Validate:       validate,
	}
}

func (s *RoleServiceImpl) FindAll() ([]response.RoleResponse, error) {
	roles, err := s.RoleRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var roleResponses []response.RoleResponse
	for _, role := range roles {
		roleResponses = append(roleResponses, response.RoleResponse{
			Id:   role.Id,
			Name: role.Name,
		})
	}
	return roleResponses, nil
}

func (s *RoleServiceImpl) FindById(roleId int) (response.RoleResponse, error) {
	role, err := s.RoleRepository.FindById(roleId)
	if err != nil {
		return response.RoleResponse{}, err
	}

	return response.RoleResponse{
		Id:   role.Id,
		Name: role.Name,
	}, nil
}

func (s *RoleServiceImpl) Create(req request.CreateRoleRequest) error {
	if err := s.Validate.Struct(req); err != nil {
		return err
	}

	role := model.Role{
		Name: req.Name,
	}
	return s.RoleRepository.Save(role)
}

func (s *RoleServiceImpl) Update(req request.UpdateRoleRequest) error {
	if err := s.Validate.Struct(req); err != nil {
		return err
	}

	role, err := s.RoleRepository.FindById(req.Id)
	if err != nil {
		return err
	}

	role.Name = req.Name
	return s.RoleRepository.Update(role)
}

func (s *RoleServiceImpl) Delete(roleId int) error {
	return s.RoleRepository.Delete(roleId)
}