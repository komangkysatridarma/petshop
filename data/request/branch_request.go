package request

type CreateBranchRequest struct {
	Name     string `validate:"required,min=1,max=150" json:"name"`
	Code     string `validate:"required,min=1,max=150" json:"code"`
	Address  string `validate:"required,min=1,max=500" json:"address"`
	Phone    string `validate:"required,min=10,max=20" json:"phone"`
	Timezone string `validate:"required,min=1,max=50" json:"timezone"`
}

type UpdateBranchRequest struct {
	Id       int    `validate:"required"`
	Name     string `validate:"required,min=1,max=150" json:"name"`
	Code     string `validate:"required,min=1,max=150" json:"code"`
	Address  string `validate:"required,min=1,max=500" json:"address"`
	Phone    string `validate:"required,min=10,max=20" json:"phone"`
	Timezone string `validate:"required,min=1,max=50" json:"timezone"`
}
