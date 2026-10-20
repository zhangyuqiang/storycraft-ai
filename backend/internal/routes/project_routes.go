package routes

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "storycraft-ai/backend/internal/dto"
    "storycraft-ai/backend/internal/service"
)

func RegisterProjectRoutes(r *gin.RouterGroup) {
    svc := service.NewProjectService()

    r.GET("/projects", func(c *gin.Context) {
        projects, err := svc.List()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"items": projects})
    })

    r.POST("/projects", func(c *gin.Context) {
        var req dto.CreateProjectRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        project, err := svc.Create(req)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, gin.H{"data": project})
    })
}
