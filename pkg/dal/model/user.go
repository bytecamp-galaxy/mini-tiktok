package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"index:idx_username,unique;type:varchar(40);not null" json:"username"`
	Password string `gorm:"type:varchar(256);not null" json:"password"`
}
