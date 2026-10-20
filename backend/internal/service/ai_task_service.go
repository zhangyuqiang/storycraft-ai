package service

import (
    "fmt"
    "time"

    "storycraft-ai/backend/internal/pkg/ai"
)

type AIJobService struct{}

func NewAIJobService() *AIJobService {
    return &AIJobService{}
}

func (s *AIJobService) BuildAssetResult(prompt, url, kind, provider, model string) ai.AssetResult {
    return ai.AssetResult{
        Prompt:     prompt,
        URL:        url,
        Kind:       kind,
        Provider:   provider,
        Model:      model,
        GeneratedAt: time.Now(),
    }
}

func (s *AIJobService) BuildStoryTask(topic, genre string) map[string]string {
    return map[string]string{
        "task":      "story_outline",
        "topic":     topic,
        "genre":     genre,
        "provider":  "llm",
        "status":    string(ai.TaskPending),
        "message":   fmt.Sprintf("Generating story outline for %s in %s", topic, genre),
    }
}
