package dto

type StoryGenerateRequest struct {
    Topic string `json:"topic" binding:"required"`
    Genre string `json:"genre" binding:"required"`
    Tone  string `json:"tone"`
}

type ShotGenerateRequest struct {
    Prompt string `json:"prompt" binding:"required"`
    Style  string `json:"style"`
}
