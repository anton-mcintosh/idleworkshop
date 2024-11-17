package main

import (
  "context"
  "log"
  "os"

  "github.com/urfave/cli/v2"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "github.com/joho/godotenv"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init()  {

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
}

func main() {
  app := &cli.App{
        Name: "Plog Boster",
        Usage: "Bost a Plog!",
        Commands: []*cli.Command{},
    }
  err2 := app.Run(os.Args)
  if err2 != nil {
    log.Fatal(err2)
  }
}
