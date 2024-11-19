package utils

func dbConnect() string {
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
  return client
}
