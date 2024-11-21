package models

import (
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPost struct {
  ID      primitive.NewObjectID()     `json:"id" bson:"id"`
  Date    primitive.NewDateTimeFromTime(time.Now()) `json:"date" bson:"date"`
  Title   string `json:"title" bson:"title"`
  Content string `json:"content" bson:"content"`
}
