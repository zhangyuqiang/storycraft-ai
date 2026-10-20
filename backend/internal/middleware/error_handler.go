package middleware

import "github.com/gin-gonic/gin"

func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if r := recover(); r != nil {
                c.JSON(500, gin.H{"error": "internal_server_error"})
                c.Abort()
            }
        }()
        c.Next()
    }
}
