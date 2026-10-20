package service

import (
	"fmt"
	"time"

	"storycraft-ai/backend/config"
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
	prompt := fmt.Sprintf("You are a creative screenwriter. Create a story outline for a %s story themed around %s. Return JSON with title, synopsis, characters, and chapter outline.", genre, topic)
	text, err := s.LLM.GenerateText(prompt)
	if err != nil {
		return map[string]any{"error": err.Error()}
	}
	return map[string]any{
		"title": fmt.Sprintf("%s - %s", genre, topic),
		"outline": text,
		"generated_at": time.Now().UTC().Format(time.RFC3339),
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
		"generated_at": time.Now().UTC().Format(time.RFC3339),
	}
}
