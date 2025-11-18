package request

type CreateCategoryRequest struct {
	Name string `validate:"required,min=1,max=100" json:"name"`
}

type UpdateCategoryRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,min=1,max=100" json:"name"`
}