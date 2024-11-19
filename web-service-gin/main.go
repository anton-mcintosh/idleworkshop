package main

import (
  "context"
  "encoding/json"
  "log"

  //"github.com/urfave/cli/v2"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  
  "idleworkshop/website/models"
  "idleworkshop/website/utils"
)

var collection *mongo.Collection
var ctx = context.TODO()

func main() {

  client := utils.dbConnect
  // access the collection
  collection = client.Database("blogs").Collection("posts")
  /*
  newPost := BlogPost{ID: 1, Title: "Testies!", Content: "One, two!"}

  result, err := collection.InsertOne(ctx, newPost)
  if err != nil {
    log.Fatal(err)
  }
  */

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
