package model

import "petshop/enum"

type User struct {
	Id           int           `gorm:"type:int;primary_key"`
	Name         string        `gorm:"type:varchar(255)"`
	Email        string        `gorm:"type:varchar(255)"`
	Password     string        `gorm:"type:varchar(255)"`
	Role         enum.UserRole `gorm:"type:enum('Admin', 'Owner', 'Staff')"`
	Phone_number string        `gorm:"type:varchar(20)"`
	Branch_id    int           `gorm:"type:int;foreign_key"`
}
