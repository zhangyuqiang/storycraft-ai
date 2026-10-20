package dto

type AIJobRequest struct {
    ProjectID int64  `json:"project_id"`
    UserID    int64  `json:"user_id"`
    TaskType  string `json:"task_type" binding:"required"`
    Provider  string `json:"provider" binding:"required"`
    ModelName string `json:"model_name"`
    Prompt    string `json:"prompt" binding:"required"`
    InputJSON string `json:"input_json"`
}

type AIJobResponse struct {
    ID        int64  `json:"id"`
    TaskType  string `json:"task_type"`
    Status    string `json:"status"`
    ResultURL string `json:"result_url,omitempty"`
}
