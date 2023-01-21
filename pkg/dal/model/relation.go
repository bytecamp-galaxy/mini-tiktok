package model

// Relation TODO
type Relation struct {
	ID       int64 `gorm:"primaryKey;" json:"id"`
	User     User  `gorm:"foreignKey:UserID;references:ID;" json:"user"`
	UserID   int64 `gorm:"index:idx_user_id;" json:"user_id"`
	ToUser   User  `gorm:"foreignKey:ToUserID;references:ID;" json:"to_user"`
	ToUserID int64 `gorm:"index:idx_to_user_id;" json:"to_user_id"`
}
