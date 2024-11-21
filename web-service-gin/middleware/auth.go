package middleware

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "idleworkshop/website/utils"
)

func AuthMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    token := c.GetHeader("Authorization")
    if token != utils.GetVar("API_KEY") {
      c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
      c.Abort()
      return
    }
    c.Next()
      }
}
