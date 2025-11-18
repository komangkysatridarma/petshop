package model

type Role struct {
	Id   int    `gorm:"type:int;primary_key"`
	Name string `gorm:"type:varchar(15)"`
}