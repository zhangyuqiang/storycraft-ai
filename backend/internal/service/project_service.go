package service

import (
    "storycraft-ai/backend/internal/dto"
    "storycraft-ai/backend/internal/model"
    "storycraft-ai/backend/internal/repository"
)

type ProjectService struct {
    repo *repository.ProjectRepository
}

func NewProjectService() *ProjectService {
    return &ProjectService{repo: repository.NewProjectRepository()}
}

func (s *ProjectService) List() ([]model.Project, error) {
    return s.repo.List()
}

func (s *ProjectService) Create(req dto.CreateProjectRequest) (model.Project, error) {
    project := model.Project{
        OwnerID:     1,
        Title:       req.Title,
        Description: req.Description,
        Genre:       req.Genre,
        Visibility:  req.Visibility,
        Status:      "draft",
    }
    return s.repo.Create(project)
}
