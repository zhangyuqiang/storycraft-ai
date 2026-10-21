package model

import (
    "time"
)

type AITask struct {
    ID          int64      `gorm:"primaryKey;autoIncrement" json:"id"`
    ProjectID   int64      `gorm:"not null;index" json:"project_id"`
    UserID      int64      `gorm:"not null;index" json:"user_id"`
    TaskType    string     `gorm:"size:50;not null" json:"task_type"`
    Provider    string     `gorm:"size:100;not null" json:"provider"`
    ModelName   string     `gorm:"size:255" json:"model_name,omitempty"`
    Status      string     `gorm:"size:32;default:pending" json:"status"`
    Prompt      string     `gorm:"type:text;not null" json:"prompt"`
    InputJSON   string     `gorm:"type:json" json:"input_json,omitempty"`
    OutputJSON  string     `gorm:"type:json" json:"output_json,omitempty"`
    ResultURL   string     `gorm:"size:500" json:"result_url,omitempty"`
    ErrorString string     `gorm:"type:text" json:"error_message,omitempty"`
    Progress    int        `gorm:"default:0" json:"progress"`
    CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
    CompletedAt *time.Time `gorm:"index" json:"completed_at,omitempty"`
}

func (AITask) TableName() string {
    return "ai_tasks"
}
