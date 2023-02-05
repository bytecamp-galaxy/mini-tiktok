package model

// User Many To Many (User, Video)
type User struct {
	ID             int64  `gorm:"primaryKey;" json:"id"`
	Username       string `gorm:"uniqueIndex:idx_user_name;type:varchar(31);" json:"username"`
	Password       string `gorm:"type:varchar(63);not null;" json:"password"`
	FollowingCount int64  `gorm:"default:0;" json:"following_count"`
	FollowerCount  int64  `gorm:"default:0;" json:"follower_count"`
}
