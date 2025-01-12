package controllers

import (
  "context"
  "log"
  "net/http"
  "time"

  "github.com/gin-gonic/gin"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "go.mongodb.org/mongo-driver/bson"

  "idleworkshop/website/models"
  "idleworkshop/website/utils"
)

func CreatePost(c *gin.Context, collection *mongo.Collection, ctx context.Context) {
  var markdownData struct {
    Markdown string `json:"markdown"`
    File string `json:"file"`
  }
  if err := c.ShouldBindJSON(&markdownData); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"oopsie-1!": err.Error()})
    return
  }
  log.Println(markdownData.Markdown)

  parsedPost, err := utils.ParseMarkdown(markdownData.Markdown)

  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"oopsie-2!": err.Error()})
    return
  }
  var slug = markdownData.File

  filter := bson.M{"slug": slug}
  update := bson.M{
    "$set": bson.M{
      "slug": slug,
      "title": parsedPost.Metadata.Title,
      "tags": parsedPost.Metadata.Tags,
      "nutshell": parsedPost.Metadata.Nutshell,
      "readtime": parsedPost.Metadata.ReadTime,
      "topic": parsedPost.Metadata.Topic,
      "content": parsedPost.Content,
      "summary": parsedPost.Summary,
    },
    "$setOnInsert": bson.M{"date": primitive.NewDateTimeFromTime(time.Now())},
  }

  options := options.Update().SetUpsert(true)

  result, err := collection.UpdateOne(ctx, filter, update, options)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"oopsie!": "Error upserting post"})
    return
      }

  if result.UpsertedCount > 0 {
    c.JSON(http.StatusCreated, gin.H{"message": "Post created"})
      } else {
    c.JSON(http.StatusOK, gin.H{"message": "Post updated"})
      }

}

func GetPosts(c *gin.Context, collection *mongo.Collection, ctx context.Context) {
  
  // define filter and option
  filter := bson.D{}
  findOptions := options.Find()
  findOptions.SetLimit(20)
  findOptions.SetSort(bson.D{{"date", -1}})

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
