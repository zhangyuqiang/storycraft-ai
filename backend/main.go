package main

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "storycraft-ai/backend/config"
    "storycraft-ai/backend/internal/api/v1"
)

type Server struct {
    Router *gin.Engine
}

func NewServer() *Server {
    cfg := config.Load()
    _ = cfg

    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status":  "ok",
            "service": "storycraft-ai-backend",
        })
    })

    api := v1.NewRouter()
    r.Any("/api/v1/*path", func(c *gin.Context) {
        if c.Request.Method == http.MethodOptions {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }
        api.HandleContext(c)
    })

    return &Server{Router: r}
}

func main() {
    cfg := config.Load()
    server := NewServer()
    if err := server.Router.Run(":" + cfg.Port); err != nil {
        panic(err)
    }
}
