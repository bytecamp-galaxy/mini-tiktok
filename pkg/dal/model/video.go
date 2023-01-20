package model

import "gorm.io/gorm"

// Video belongs to User
type Video struct {
	gorm.Model
	Author        User   `gorm:"foreignKey:AuthorID;" json:"author"`
	AuthorID      int    `gorm:"index:idx_author_id;" json:"author_id"`
	PlayUrl       string `gorm:"type:varchar(255);not null;" json:"play_url"`
	CoverUrl      string `gorm:"type:varchar(255);" json:"cover_url"`
	FavoriteCount int    `gorm:"default:0;" json:"favorite_count"`
	CommentCount  int    `gorm:"default:0;" json:"comment_count"`
	Title         string `gorm:"type:varchar(63);not null;" json:"title"`
}
