package main

import (
  "context"
  "log"
  "os"

  //"github.com/urfave/cli/v2"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "github.com/joho/godotenv"
)

type BlogPost struct {
  ID      int    `json:"id"`
  Title   string `json:"title"`
  Content string `json:"content"`
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
  collection = client.Database("blogs").Collection("posts")
  newPost := BlogPost{ID: 1, Title: "Testies!", Content: "One, two!"}

  result, err := collection.InsertOne(ctx, newPost)
  if err != nil {
    log.Fatal(err)
  }
  filter := BlogPost{Title: "Testies!"}
  err = collection.FindOne(ctx, filter).Decode(&newPost)

  if err != nil {
    log.Fatal(err)
      }
  log.println(result)
  log.Println(newPost)

}
