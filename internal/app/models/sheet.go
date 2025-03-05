package models

type Page struct {
	ID          uint      `gorm:"primary_key;auto_increment" json:"id"`
	Source      string    `gorm:"type:text;not null" json:"source"`
	PageNumber  uint      `json:"page_number"`
	CharacterID uint      `json:"character_id"`
	Character   Character `gorm:"foreignKey:CharacterID" json:"-"`
	DiaryID     uint      `json:"diary_id"`
	Diary       Diary     `gorm:"foreignKey:DiaryID" json:"-"`
}
