package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	Id          int             `gorm:"type:int;primary_key"`
	Name        string          `gorm:"type:varchar(100)"`
	Code        string          `gorm:"type:varchar(30);uniqueIndex"`
	Price       decimal.Decimal `gorm:"type:decimal(10,2)"`
	Category_id int             `gorm:"type:int;foreign_key"`
	Is_service  bool            `gorm:"type:boolean"`
	Track_batch bool            `gorm:"type:boolean"`
	Created_at  time.Time       `gorm:"type:timestamp"`
}
