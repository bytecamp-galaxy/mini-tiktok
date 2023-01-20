package model

import "gorm.io/gorm"

// User Many To Many (User, Video)
type User struct {
	gorm.Model
	Username       string  `gorm:"uniqueIndex:index_user_name;type:varchar(31);" json:"username"`
	Password       string  `gorm:"type:varchar(63);not null;" json:"password"`
	FavoriteVideos []Video `gorm:"many2many:user_favorite_videos;" json:"favorite_videos"`
	FollowingCount int     `gorm:"default:0;" json:"following_count"`
	FollowerCount  int     `gorm:"default:0;" json:"follower_count"`
}
