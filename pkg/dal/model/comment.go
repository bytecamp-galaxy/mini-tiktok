package model

// Comment belongs to Video and User
type Comment struct {
	ID      int64  `gorm:"primaryKey;" json:"id"`
	Video   Video  `gorm:"foreignKey:VideoID;references:ID;" json:"video"`
	VideoID int64  `gorm:"index:idx_video_id,sort:desc;" json:"video_id"`
	User    User   `gorm:"foreignKey:UserID;references:ID;" json:"user"`
	UserID  int64  `gorm:"index:idx_user_id;" json:"user_id"`
	Content string `gorm:"type:varchar(255);not null;" json:"content"`
}
