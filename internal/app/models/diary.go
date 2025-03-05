package models

type Diary struct {
	ID     uint   `gorm:"primary_key;auto_increment" json:"id"`
	Title  string `gorm:"size:255;not null" json:"title"`
	Pages  uint   `gorm:"default:0" json:"pages"`
	Rate   uint   `gorm:"null" json:"rate"`
	UserID uint   `gorm:"not null" json:"user_id"`
	User   User   `gorm:"foreignkey:UserID" json:"-"`
}
