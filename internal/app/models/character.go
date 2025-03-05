package models

type Character struct {
	ID      uint   `gorm:"primary_key;auto_increment" json:"id"`
	Title   string `gorm:"size:255;not null" json:"title"`
	DiaryID uint   `gorm:"not null" json:"diary_id"`
	Diary   Diary  `gorm:"foreignkey:DiaryID" json:"-"`
	UserID  uint   `gorm:"not null" json:"user_id"`
	User    User   `gorm:"foreignkey:UserID" json:"-"`
}
