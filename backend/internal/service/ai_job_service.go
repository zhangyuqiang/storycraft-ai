package service

import (
    "encoding/json"
    "fmt"
    "time"

    "storycraft-ai/backend/internal/dto"
    "storycraft-ai/backend/internal/model"
    "storycraft-ai/backend/internal/pkg/ai"
    "storycraft-ai/backend/internal/repository"
)

type AIJobService struct {
    repo *repository.AIJobRepository
}

func NewAIJobService() *AIJobService {
    return &AIJobService{repo: repository.NewAIJobRepository()}
}

func (s *AIJobService) Create(req dto.AIJobRequest) (model.AITask, error) {
    job := model.AITask{
        ProjectID: req.ProjectID,
        UserID:    req.UserID,
        TaskType:  req.TaskType,
        Provider:  req.Provider,
        ModelName: req.ModelName,
        Prompt:    req.Prompt,
        InputJSON: req.InputJSON,
        Status:    "pending",
        Progress:  0,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
    return s.repo.Create(job)
}

func (s *AIJobService) SaveResult(taskID int64, kind, provider, modelName, resultURL string, metadata map[string]any) (model.AIResult, error) {
    payload, _ := json.Marshal(metadata)
    result := model.AIResult{
        TaskID:    taskID,
        Kind:      kind,
        Provider:  provider,
        ModelName: modelName,
        ResultURL: resultURL,
        Metadata:  string(payload),
        CreatedAt: time.Now(),
    }
    return s.repo.SaveResult(result)
}

func (s *AIJobService) BuildStatus(taskType, status string) map[string]any {
    return map[string]any{
        "task_type": taskType,
        "status":    status,
        "message":   fmt.Sprintf("AI task %s is %s", taskType, status),
        "updated_at": time.Now().UTC().Format(time.RFC3339),
    }
}

func (s *AIJobService) BuildAssetResult(prompt, url, kind, provider, model string) ai.AssetResult {
    return ai.AssetResult{
        Prompt:     prompt,
        URL:        url,
        Kind:       kind,
        Provider:   provider,
        Model:      model,
        GeneratedAt: time.Now(),
    }
}

func (s *AIJobService) BuildStoryTask(topic, genre string) map[string]string {
    return map[string]string{
        "task":     "story_outline",
        "topic":    topic,
        "genre":    genre,
        "provider": "llm",
        "status":   string(ai.TaskPending),
        "message":  fmt.Sprintf("Generating story outline for %s in %s", topic, genre),
    }
}
