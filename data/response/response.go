package response

import "petshop/enum"

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UserResponse struct {
	Id           int           `json:"id"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Password     string        `json:"password"`
	Role         enum.UserRole `json:"role"`
	Phone_number string        `json:"phone_number"`
	Branch_id    int           `json:"branch_id"`
}

type BranchResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Timezone string `json:"timezone"`
}
