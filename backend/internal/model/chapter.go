package model

import "time"

type Chapter struct {
    ID        int64     `json:"id"`
    ProjectID int64     `json:"project_id"`
    Title     string    `json:"title"`
    Summary   string    `json:"summary,omitempty"`
    SortIndex int       `json:"sort_index"`
    Status    string    `json:"status"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func (Chapter) TableName() string {
    return "chapters"
}
