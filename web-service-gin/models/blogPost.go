package models

import (
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPost struct {
  ID      primitive `json:"id" bson:"id"`
  Date    primitive `json:"date" bson:"date"`
  Title   string `json:"title" bson:"title"`
  Content string `json:"content" bson:"content"`
}
