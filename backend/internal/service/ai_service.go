package service

import (
    "fmt"
    "time"

    "storycraft-ai/backend/config"
    "storycraft-ai/backend/internal/pkg/ai"
    "storycraft-ai/backend/internal/pkg/llm"
    "storycraft-ai/backend/internal/pkg/seedance"
)

type AIService struct {
    LLM      *llm.Client
    Seedance *seedance.Client
}

func NewAIService() *AIService {
    cfg := config.Load()
    return &AIService{
        LLM:      llm.NewClient(cfg.LLMBaseURL, cfg.LLMAPIKey),
        Seedance: seedance.NewClient(cfg.SeedanceBaseURL, cfg.SeedanceAPIKey),
    }
}

func (s *AIService) GenerateStory(topic, genre string) map[string]any {
    prompt := fmt.Sprintf("You are a creative screenwriter. Return a JSON object with title, synopsis, characters, chapters, and theme. Story topic: %s. Genre: %s.", topic, genre)
    var outline ai.StoryOutline
    if err := s.LLM.GenerateJSON(prompt, &outline); err != nil {
        return map[string]any{"error": err.Error()}
    }
    if outline.GeneratedAt == "" {
        outline.GeneratedAt = time.Now().UTC().Format(time.RFC3339)
    }
    return map[string]any{
        "title": outline.Title,
        "genre": outline.Genre,
        "outline": outline,
        "generated_at": outline.GeneratedAt,
        "provider": "llm-generic",
    }
}

func (s *AIService) GenerateShot(prompt string) map[string]any {
    imageURL, err := s.Seedance.GenerateImage(prompt)
    if err != nil {
        return map[string]any{"error": err.Error()}
    }
    return map[string]any{
        "prompt": prompt,
        "image_url": imageURL,
        "kind": "image",
        "provider": "seedance",
        "generated_at": time.Now().UTC().Format(time.RFC3339),
    }
}

func (s *AIService) GenerateVideo(prompt string) map[string]any {
    videoURL, err := s.Seedance.GenerateVideo(prompt)
    if err != nil {
        return map[string]any{"error": err.Error()}
    }
    return map[string]any{
        "prompt": prompt,
        "video_url": videoURL,
        "kind": "video",
        "provider": "seedance",
        "generated_at": time.Now().UTC().Format(time.RFC3339),
    }
}
