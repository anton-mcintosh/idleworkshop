package utils

import (
  "github.com/joho/godotenv"
)

  err := godotenv.Load(".env")

  func GetVar(req string) (requestedVar string) {
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  requestedVar = os.Getenv(req)
  return
}
