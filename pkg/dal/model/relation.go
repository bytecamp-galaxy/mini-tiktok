package model

import "gorm.io/gorm"

// Relation TODO
type Relation struct {
	gorm.Model
	User     User `gorm:"foreignKey:UserID;"`
	UserID   int  `gorm:"index:idx_rel;"`
	ToUser   User `gorm:"foreignKey:ToUserID;"`
	ToUserID int  `gorm:"index:idx_rel;"`
}
