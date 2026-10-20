package model

import "time"

type AIJob struct {
    ID          int64     `json:"id"`
    ProjectID   int64     `json:"project_id"`
    UserID      int64     `json:"user_id"`
    TaskType    string    `json:"task_type"`
    Provider    string    `json:"provider"`
    ModelName   string    `json:"model_name,omitempty"`
    Status      string    `json:"status"`
    Prompt      string    `json:"prompt"`
    InputJSON   string    `json:"input_json,omitempty"`
    OutputJSON  string    `json:"output_json,omitempty"`
    ResultURL   string    `json:"result_url,omitempty"`
    ErrorString string    `json:"error_message,omitempty"`
    Progress    int       `json:"progress"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    CompletedAt time.Time `json:"completed_at,omitempty"`
}

func (AIJob) TableName() string {
    return "ai_tasks"
}
