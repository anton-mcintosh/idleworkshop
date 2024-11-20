package controllers

import (
  "context"
  "log"
  "encoding/json"

  "github.com/gin-gonic/gin"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/bson"

  "idleworkshop/website/models"
)
func GetPosts(c *gin.Context, collection *mongo.Collection, ctx context.Context) {
  
  // define filter and option
  filter := bson.D{}
  findOptions := options.Find()
  findOptions.SetLimit(5)

  // decode results
  // NTFS: To decode a single object, use the decode() method. For multiple documents, need to iterate over the cursor and decode each.
  var allPosts []models.BlogPost
  cursor, err := collection.Find(ctx, filter, findOptions)
  if err != nil {
    log.Fatal(err)
      }
  defer cursor.Close(ctx) // I need to read about this.

  //Decode the results
  for cursor.Next(ctx) {
    var post models.BlogPost
    err := cursor.Decode(&post)
    if err != nil {
      log.Println("Error decoding post: ", err)
      continue
    }
    allPosts = append(allPosts, post)
  }
  if err := cursor.Err(); err != nil {
    log.Fatal(err)
  }
  for _, post := range allPosts {
    postJSON, _ := json.Marshal(post)
    log.Println(string(postJSON))
  }
}
