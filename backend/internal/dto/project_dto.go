package dto

type CreateProjectRequest struct {
    Title       string `json:"title" binding:"required"`
    Description string `json:"description"`
    Genre       string `json:"genre"`
    Visibility  string `json:"visibility"`
}

type UpdateProjectRequest struct {
    Title       string `json:"title"`
    Description string `json:"description"`
    Genre       string `json:"genre"`
    Visibility  string `json:"visibility"`
    Status      string `json:"status"`
}

