package main

import (
  "context"

  "github.com/gin-gonic/gin"
  "go.mongodb.org/mongo-driver/mongo"
  "github.com/gin-contrib/cors"

  
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
  config := cors.Config{
    AllowOrigins: []string{"*", "75.172.97.44"},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Authorization", "Content-Type", "Origin"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
  }
  router.OPTIONS("/*path", func(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
    c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, Origin")
    c.Status(204) // No Content
})
  router.Use(cors.New(config))
  router.GET("/get-posts", func(c *gin.Context) {(controllers.GetPosts(c, collection, ctx))})
  router.POST("/posts", middleware.AuthMiddleware(), func (c *gin.Context) {(controllers.CreatePost(c, collection, ctx))})
  router.Run("0.0.0.0:8080")
  /*
  newPost := BlogPost{ID: 1, Title: "Testies!", Content: "One, two!"}

  result, err := collection.InsertOne(ctx, newPost)
  if err != nil {
    log.Fatal(err)
  }
  */

}
