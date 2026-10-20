package model

import "time"

type AIResult struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	TaskID    int64     `gorm:"not null;index" json:"task_id"`
	Kind      string    `gorm:"size:32;not null" json:"kind"`
	Provider  string    `gorm:"size:100;not null" json:"provider"`
	ModelName string    `gorm:"size:255" json:"model_name,omitempty"`
	ResultURL string    `gorm:"size:500" json:"result_url,omitempty"`
	Metadata  string    `gorm:"type:json" json:"metadata,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (AIResult) TableName() string {
	return "ai_results"
}
