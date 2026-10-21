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

func (s *ProjectService) GetByID(id int64) (model.Project, error) {
    return s.repo.GetByID(id)
}

func (s *ProjectService) Create(req dto.CreateProjectRequest) (model.Project, error) {
    if req.Title == "" {
        return model.Project{}, nil
    }

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

func (s *ProjectService) Update(id int64, req dto.UpdateProjectRequest) (model.Project, error) {
    project, err := s.repo.GetByID(id)
    if err != nil {
        return model.Project{}, err
    }

    if req.Title != "" {
        project.Title = req.Title
    }
    if req.Description != "" {
        project.Description = req.Description
    }
    if req.Genre != "" {
        project.Genre = req.Genre
    }
    if req.Visibility != "" {
        project.Visibility = req.Visibility
    }
    if req.Status != "" {
        project.Status = req.Status
    }

    return s.repo.Update(project)
}
