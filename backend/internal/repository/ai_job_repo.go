package repository

import (
	"storycraft-ai/backend/internal/db"
	"storycraft-ai/backend/internal/model"
)

type AIJobRepository struct{}

func NewAIJobRepository() *AIJobRepository {
	return &AIJobRepository{}
}

func (r *AIJobRepository) Create(job model.AITask) (model.AITask, error) {
	if db.DB == nil {
		return job, nil
	}
	if err := db.DB.Create(&job).Error; err != nil {
		return model.AITask{}, err
	}
	return job, nil
}

func (r *AIJobRepository) SaveResult(result model.AIResult) (model.AIResult, error) {
	if db.DB == nil {
		return result, nil
	}
	if err := db.DB.Create(&result).Error; err != nil {
		return model.AIResult{}, err
	}
	return result, nil
}
