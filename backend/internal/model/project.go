package model

import "time"

type Project struct {
    ID          int64     `json:"id"`
    OwnerID     int64     `json:"owner_id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    CoverURL    string    `json:"cover_url,omitempty"`
    Genre       string    `json:"genre,omitempty"`
    Visibility  string    `json:"visibility"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

func (Project) TableName() string {
    return "projects"
}

