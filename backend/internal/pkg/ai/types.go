package ai

import "time"

type TaskStatus string

const (
    TaskPending   TaskStatus = "pending"
    TaskRunning   TaskStatus = "running"
    TaskSuccess   TaskStatus = "success"
    TaskFailed    TaskStatus = "failed"
)

type StoryOutline struct {
    Title       string   `json:"title"`
    Synopsis    string   `json:"synopsis"`
    Genre       string   `json:"genre"`
    Theme       string   `json:"theme"`
    Characters  []string `json:"characters"`
    Chapters    []string `json:"chapters"`
    GeneratedAt string   `json:"generated_at"`
}

type AssetResult struct {
    Prompt     string    `json:"prompt"`
    URL        string    `json:"url"`
    Kind       string    `json:"kind"`
    Provider   string    `json:"provider"`
    Model      string    `json:"model"`
    GeneratedAt time.Time `json:"generated_at"`
}

type VideoRequest struct {
    Prompt string `json:"prompt"`
    Model  string `json:"model,omitempty"`
    Width  int    `json:"width,omitempty"`
    Height int    `json:"height,omitempty"`
}
