package request

import "petshop/enum"

type CreateUserRequest struct {
	Name         string        `validate:"required,min=1,max=200" json:"name"`
	Email        string        `validate:"required,email,min=5,max=200" json:"email"`
	Password     string        `validate:"required,min=8,max=200" json:"password"`
	Role         enum.UserRole `validate:"required,oneof=Admin Owner Staff" json:"role"`
	Phone_number string        `validate:"required,min=10,max=20" json:"phone_number"`
	Branch_id    int           `validate:"min=1" json:"branch_id"`
}

type UpdateUserRequest struct {
	Id           int           `validate:"required"`
	Name         string        `validate:"required,min=1,max=200" json:"name"`
	Email        string        `validate:"required,email,min=5,max=200" json:"email"`
	Password     string        `validate:"required,min=8,max200" json:"password"`
	Role         enum.UserRole `validate:"required,oneof=Admin Owner Staff" json:"role"`
	Phone_number string        `validate:"required,min=10,max=20" json:"phone_number"`
	Branch_id    int           `validate:"min=1" json:"branch_id"`
}
