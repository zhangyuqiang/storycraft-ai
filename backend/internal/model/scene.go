package model

import "time"

type Scene struct {
    ID          int64     `json:"id"`
    ProjectID   int64     `json:"project_id"`
    ChapterID   int64     `json:"chapter_id"`
    Title       string    `json:"title"`
    SceneNumber int       `json:"scene_number"`
    Location    string    `json:"location,omitempty"`
    Mood        string    `json:"mood,omitempty"`
    Summary     string    `json:"summary,omitempty"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

func (Scene) TableName() string {
    return "scenes"
}
