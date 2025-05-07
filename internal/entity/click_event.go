package entity

import (
	"time"
)

type ClickEvent struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	LinkID    uint      `gorm:"not null;index"` // foreign key field
	Link      *Link     `gorm:"constraint:OnDelete:CASCADE;"`
	ClickedAt time.Time `gorm:"autoCreateTime" json:"clicked_at"` // Timestamp of the click
	IPAddress string    `gorm:"size:45" json:"ip_address"`        // IPv4 or IPv6
	UserAgent string    `gorm:"type:text" json:"user_agent"`      // Raw user agent string
	Country   string    `gorm:"size:100" json:"country"`          // Optional, from GeoIP
	Referrer  string    `gorm:"type:text" json:"referrer"`        // Optional
	Platform  string    `gorm:"size:100" json:"platform"`         // Parsed OS
	Browser   string    `gorm:"size:100" json:"browser"`          // Parsed browser
}
