package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"index:index_username,unique;type:varchar(32);not null" json:"username"`
	Password string `gorm:"type:varchar(64);not null" json:"password"`
}
