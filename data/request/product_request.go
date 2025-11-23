package request

import "github.com/shopspring/decimal"

type CreateProductRequest struct {
	Name        string          `validate:"required,min=3,max=100" json:"name"`
	Code        string          `validate:"required,min=2,max=30;uniqueIndex" json:"code"`
	Price       decimal.Decimal `validate:"required" json:"price"`
	Category_id int             `validate:"required" json:"category_id"`
	Is_service  bool            `json:"is_service"`
	Track_batch bool            `json:"track_batch"`
}

type UpdateProductRequest struct {
	Id          int             `validate:"required" json:"id"`
	Name        string          `validate:"required,min=3,max=100" json:"name"`
	Code        string          `validate:"required,min=2,max=30;uniqueIndex" json:"code"`
	Price       decimal.Decimal `validate:"required" json:"price"`
	Category_id int             `validate:"required" json:"category_id"`
	Is_service  bool            `json:"is_service"`
	Track_batch bool            `json:"track_batch"`
}
