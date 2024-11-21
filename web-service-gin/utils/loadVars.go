package utils

import (
  "log"
  "os"
  "github.com/joho/godotenv"
)

  func GetVar(req string) (requestedVar string) {
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  requestedVar = os.Getenv(req)
  return
}
