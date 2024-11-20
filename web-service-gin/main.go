package main

import (
  "context"

  "github.com/gin-gonic/gin"
  "go.mongodb.org/mongo-driver/mongo"
  
  "idleworkshop/website/utils"
  "idleworkshop/website/controllers"
)

var collection *mongo.Collection
var ctx = context.TODO()

func main() {

  client := utils.DBConnect()
  // access the collection
  collection = client.Database("blogs").Collection("posts")
  router := gin.Default()
  router.GET("/posts", func (c *gin.context){controllers.GetPosts(collection, ctx)})
  router.Run("0.0.0.0:8080")
  /*
  newPost := BlogPost{ID: 1, Title: "Testies!", Content: "One, two!"}

  result, err := collection.InsertOne(ctx, newPost)
  if err != nil {
    log.Fatal(err)
  }
  */

}
