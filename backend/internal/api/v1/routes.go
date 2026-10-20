package v1

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "storycraft-ai/backend/internal/routes"
    "storycraft-ai/backend/internal/service"
)

func NewRouter() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger(), gin.Recovery())

    api := r.Group("/api/v1")

    api.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "storycraft-ai"})
    })

    routes.RegisterProjectRoutes(api)
    routes.RegisterAIJobRoutes(api)

    api.POST("/ai/story/generate", func(c *gin.Context) {
        var req struct {
            Topic string `json:"topic"`
            Genre string `json:"genre"`
            Tone  string `json:"tone"`
        }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        result := service.NewAIService().GenerateStory(req.Topic, req.Genre)
        c.JSON(http.StatusOK, gin.H{"data": result})
    })

    api.POST("/ai/shot/generate", func(c *gin.Context) {
        var req struct {
            Prompt string `json:"prompt"`
            Style  string `json:"style"`
        }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        result := service.NewAIService().GenerateShot(req.Prompt)
        c.JSON(http.StatusOK, gin.H{"data": result})
    })

    api.POST("/ai/video/generate", func(c *gin.Context) {
        var req struct {
            Prompt string `json:"prompt"`
        }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        result := service.NewAIService().GenerateVideo(req.Prompt)
        c.JSON(http.StatusOK, gin.H{"data": result})
    })

    return r
}
