package main

import (
  "context"
  "encoding/json"
  "log"
  "os"

  //"github.com/urfave/cli/v2"
  "go.mongo.db/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "github.com/joho/godotenv"
)

type BlogPost struct {
  ID      int    `json:"id" bson:"id"`
  Title   string `json:"title" bson:"title"`
  Content string `json:"content" bson:"content"`
}

var collection *mongo.Collection
var ctx = context.TODO()

func main() {
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  mongoURI := os.Getenv("MONGO_URI")

  clientOptions := options.Client().ApplyURI(mongoURI)
  client, err := mongo.Connect(ctx, clientOptions)
  if err != nil {
    log.Fatal(err)
      }
  err = client.Ping(ctx, nil)
  if err != nil {
    log.Fatal(err)
      }
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
  var allPosts []BlogPost
  cursor, err := collection.Find(ctx, filter, findOptions)
  if err != nil {
    log.Fatal(err)
      }
  defer cursor.Close(ctx) // I need to read about this.

  //Decode the results
  for cursor.Next(ctx) {
    var post BlogPost
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
