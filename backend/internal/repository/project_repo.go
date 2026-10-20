package repository

import (
    "storycraft-ai/backend/internal/model"
)

type ProjectRepository struct{}

func NewProjectRepository() *ProjectRepository {
    return &ProjectRepository{}
}

func (r *ProjectRepository) List() ([]model.Project, error) {
    return []model.Project{
        {
            ID:          1,
            OwnerID:     1,
            Title:       "Moonfall City",
            Description: "A city suspended above a drowned world.",
            Genre:       "fantasy",
            Visibility:  "private",
            Status:      "draft",
        },
        {
            ID:          2,
            OwnerID:     1,
            Title:       "Glass Horizon",
            Description: "A broken colony faced with impossible memory.",
            Genre:       "scifi",
            Visibility:  "team",
            Status:      "in_progress",
        },
    }, nil
}

func (r *ProjectRepository) Create(p model.Project) (model.Project, error) {
    return p, nil
}
