package model

import "time"

type Project struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	OwnerID     int64     `gorm:"not null;index" json:"owner_id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	CoverURL    string    `gorm:"size:500" json:"cover_url,omitempty"`
	Genre       string    `gorm:"size:100" json:"genre,omitempty"`
	Visibility  string    `gorm:"size:32;default:private" json:"visibility"`
	Status      string    `gorm:"size:32;default:draft" json:"status"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (Project) TableName() string {
	return "projects"
}
