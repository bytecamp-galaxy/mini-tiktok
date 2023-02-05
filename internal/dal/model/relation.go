package model

// FollowRelation
// unique multi-column index for user id and to user id
// extra single index for to user id
type FollowRelation struct {
	ID       int64 `gorm:"primaryKey;" json:"id"`
	User     User  `gorm:"foreignKey:UserID;references:ID;" json:"user"`
	UserID   int64 `gorm:"uniqueIndex:idx_rel;" json:"user_id"`
	ToUser   User  `gorm:"foreignKey:ToUserID;references:ID;" json:"to_user"`
	ToUserID int64 `gorm:"uniqueIndex:idx_rel;index:idx_to_user_id;" json:"to_user_id"`
}

// FavoriteRelation
// unique multi-column index for user id and video id
type FavoriteRelation struct {
	ID      int64 `gorm:"primaryKey;" json:"id"`
	User    User  `gorm:"foreignKey:UserID;references:ID;" json:"user"`
	UserID  int64 `gorm:"uniqueIndex:idx_rel;" json:"user_id"`
	Video   Video `gorm:"foreignKey:VideoID;references:ID;" json:"video"`
	VideoID int64 `gorm:"uniqueIndex:idx_rel;" json:"video_id"`
}
