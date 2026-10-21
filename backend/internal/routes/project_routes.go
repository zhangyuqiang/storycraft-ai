package routes

import (
    "net/http"
    "strconv"

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

    r.GET("/projects/:id", func(c *gin.Context) {
        id, err := strconv.ParseInt(c.Param("id"), 10, 64)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project id"})
            return
        }

        project, err := svc.GetByID(id)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "project not found"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"data": project})
    })

    r.POST("/projects", func(c *gin.Context) {
        var req dto.CreateProjectRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if strings.TrimSpace(req.Title) == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "project title is required"})
            return
        }
        project, err := svc.Create(req)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, gin.H{"data": project})
    })

    r.PUT("/projects/:id", func(c *gin.Context) {
        id, err := strconv.ParseInt(c.Param("id"), 10, 64)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project id"})
            return
        }

        var req dto.UpdateProjectRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        project, err := svc.Update(id, req)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"data": project})
    })
}
