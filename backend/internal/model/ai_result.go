package model

import "time"

type AIResult struct {
    ID          int64     `json:"id"`
    TaskID      int64     `json:"task_id"`
    Kind        string    `json:"kind"`
    Provider    string    `json:"provider"`
    ModelName   string    `json:"model_name,omitempty"`
    ResultURL   string    `json:"result_url,omitempty"`
    Metadata    string    `json:"metadata,omitempty"`
    CreatedAt   time.Time `json:"created_at"`
}

func (AIResult) TableName() string {
    return "ai_results"
}
