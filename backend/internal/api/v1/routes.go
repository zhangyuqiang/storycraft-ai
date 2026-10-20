package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"storycraft-ai/backend/internal/service"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "storycraft-ai"})
	})

	r.GET("/api/v1/projects", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"items": []map[string]any{}})
	})

	r.POST("/api/v1/projects", func(c *gin.Context) {
		var body map[string]any
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "project created", "data": body})
	})

	r.POST("/api/v1/ai/story/generate", func(c *gin.Context) {
		var req struct {
			Topic string `json:"topic"`
			Genre string `json:"genre"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := service.NewAIService().GenerateStory(req.Topic, req.Genre)
		c.JSON(http.StatusOK, gin.H{"data": result})
	})

	r.POST("/api/v1/ai/shot/generate", func(c *gin.Context) {
		var req struct {
			Prompt string `json:"prompt"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := service.NewAIService().GenerateShot(req.Prompt)
		c.JSON(http.StatusOK, gin.H{"data": result})
	})

	return r
}
