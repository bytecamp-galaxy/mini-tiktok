package model

import "gorm.io/gorm"

// Comment belongs to Video and User
type Comment struct {
	gorm.Model
	Video   Video  `gorm:"foreignKey:VideoID;" json:"video"`
	VideoID int    `gorm:"index:idx_video_id;" json:"video_id"`
	User    User   `gorm:"foreignKey:UserID;" json:"user"`
	UserID  int    `gorm:"index:idx_user_id;" json:"user_id"`
	Content string `gorm:"type:varchar(255);not null;" json:"content"`
}
