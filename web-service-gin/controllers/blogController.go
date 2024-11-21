package controllers

import (
  "context"
  "log"
  "net/http"

  "github.com/gin-gonic/gin"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/bson"

  "idleworkshop/website/models"
)

func CreatePost(c *gin.Context, collection *mongo.Collection, ctx context.Context) {
  var newPost models.BlogPost
  if err := c.ShouldBindJSON(&newPost); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"oopsie!": err.Error()})
    return
  }

  newPost.ID = primitive.NewObjectID()
  newPost.Date = time.Now().String()

  c.BindJSON(&post)
  _, err := collection.InsertOne(ctx, newPost)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"oopsie!": err.Error()})
  }
  c.JSON(http.StatusCreated, gin.H{"message": "Post created!"})
}

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
  c.JSON(200, gin.H{"posts": allPosts})
}
