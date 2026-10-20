package model

import "time"

type Shot struct {
    ID          int64     `json:"id"`
    ProjectID   int64     `json:"project_id"`
    ChapterID   int64     `json:"chapter_id"`
    SceneID     int64     `json:"scene_id"`
    ShotNumber  int       `json:"shot_number"`
    Title       string    `json:"title,omitempty"`
    Description string    `json:"description,omitempty"`
    CameraInfo  string    `json:"camera_info,omitempty"`
    Action      string    `json:"action,omitempty"`
    Dialogue    string    `json:"dialogue,omitempty"`
    VisualPrompt string  `json:"visual_prompt,omitempty"`
    SourceType  string    `json:"source_type"`
    AIJobID     int64     `json:"ai_job_id,omitempty"`
    CoverURL    string    `json:"cover_url,omitempty"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

func (Shot) TableName() string {
    return "shots"
}
