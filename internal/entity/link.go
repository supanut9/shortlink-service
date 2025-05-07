package entity

import "time"

type Link struct {
	ID          uint         `gorm:"primaryKey"`
	URL         string       `gorm:"type:text;not null"`
	Slug        string       `gorm:"size:50;uniqueIndex;not null"`
	Title       string       `gorm:"size:255"`
	Clicks      int          `gorm:"default:0"`
	ClickEvents []ClickEvent `gorm:"foreignKey:LinkID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
