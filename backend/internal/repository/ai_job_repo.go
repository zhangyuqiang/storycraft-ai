package repository

import (
    "storycraft-ai/backend/internal/model"
)

type AIJobRepository struct{}

func NewAIJobRepository() *AIJobRepository {
    return &AIJobRepository{}
}

func (r *AIJobRepository) Create(job model.AITask) (model.AITask, error) {
    job.ID = 1001
    job.Status = "pending"
    return job, nil
}

func (r *AIJobRepository) SaveResult(result model.AIResult) (model.AIResult, error) {
    result.ID = 2001
    return result, nil
}
