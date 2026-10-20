package routes

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "storycraft-ai/backend/internal/dto"
    "storycraft-ai/backend/internal/service"
)

func RegisterAIJobRoutes(r *gin.RouterGroup) {
    svc := service.NewAIJobService()

    r.POST("/ai/jobs", func(c *gin.Context) {
        var req dto.AIJobRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        job, err := svc.Create(req)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusAccepted, gin.H{
            "data": gin.H{
                "id":      job.ID,
                "status":  job.Status,
                "task_type": job.TaskType,
            },
        })
    })

    r.GET("/ai/jobs/:id/status", func(c *gin.Context) {
        id := c.Param("id")
        c.JSON(http.StatusOK, gin.H{
            "id": id,
            "status": "pending",
            "message": "task accepted and queued",
        })
    })
}
