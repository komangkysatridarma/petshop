package model

type Branch struct {
	Id       int    `gorm:"type:int;primary_key"`
	Name     string `gorm:"type:varchar(150)"`
	Code     string `gorm:"type:varchar(150);uniqueIndex"`
	Address  string `gorm:"type:text"`
	Phone    string `gorm:"type:varchar(50)"`
	Timezone string `gorm:"type:varchar(50)"`
}
