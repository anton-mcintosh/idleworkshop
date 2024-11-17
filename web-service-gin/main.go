package main

import (
  "context",
  "log",
  "os",

  "go.mongodb.org/mongo-driver/mongo",
  "go.mongodb.org/mongo-driver/mongo/options",
  "github.com/joho/godotenv",
)

var collection *mongo.Collection
var ctx = context.TODO()

func ()  {

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
  collection := client.Database("blogs").Collection("posts")
}
