package request

type CreateRoleRequest struct {
	Name string `validate:"required,min=1,max=15" json:"name"`
}

type UpdateRoleRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,min=1,max=15" json:"name"`
}