package main

import (
  "context"
  "net/http"

//  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "go.mongodb.org/mongo-driver/mongo"
  
  "idleworkshop/website/utils"
  "idleworkshop/website/controllers"
  "idleworkshop/website/middleware"
)

var collection *mongo.Collection
var ctx = context.TODO()

func main() {

  client := utils.DBConnect()

  // access the collection
  collection = client.Database("blogs").Collection("posts")
  router := gin.Default()
  router.Use(corsMiddleware())

  router.GET("/api/get-posts", func(c *gin.Context) {(controllers.GetPosts(c, collection, ctx))})
  router.POST("/api/posts", middleware.AuthMiddleware(), func (c *gin.Context) {(controllers.CreatePost(c, collection, ctx))})
  router.Run("0.0.0.0:8080")
}
func corsMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    if c.Request.Method == http.MethodOptions {
      c.AbortWithStatus(http.StatusOK)
      return
    }
    c.Next()
  }
}
