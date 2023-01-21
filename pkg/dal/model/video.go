package model

// Video belongs to User
type Video struct {
	ID            int64  `gorm:"primaryKey;" json:"id"`
	CreatedAt     int64  `gorm:"index:idx_created_at;autoCreateTime;" json:"created_at"` // 使用时间戳秒数填充创建时间
	Author        User   `gorm:"foreignKey:AuthorID;references:ID;" json:"author"`
	AuthorID      int64  `gorm:"index:idx_author_id;" json:"author_id"`
	PlayUrl       string `gorm:"type:varchar(255);not null;" json:"play_url"`
	CoverUrl      string `gorm:"type:varchar(255);" json:"cover_url"`
	FavoriteCount int64  `gorm:"default:0;" json:"favorite_count"`
	CommentCount  int64  `gorm:"default:0;" json:"comment_count"`
	Title         string `gorm:"type:varchar(63);not null;" json:"title"`
}
