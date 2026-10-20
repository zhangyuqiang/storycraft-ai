package repository

import (
	"storycraft-ai/backend/internal/db"
	"storycraft-ai/backend/internal/model"
)

type ProjectRepository struct{}

func NewProjectRepository() *ProjectRepository {
	return &ProjectRepository{}
}

func (r *ProjectRepository) List() ([]model.Project, error) {
	var projects []model.Project
	if db.DB == nil {
		return []model.Project{}, nil
	}
	if err := db.DB.Order("created_at desc").Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *ProjectRepository) Create(p model.Project) (model.Project, error) {
	if db.DB == nil {
		return p, nil
	}
	if err := db.DB.Create(&p).Error; err != nil {
		return model.Project{}, err
	}
	return p, nil
}
