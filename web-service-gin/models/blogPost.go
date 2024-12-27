package models

import (
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPost struct {
  ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
  Date    primitive.DateTime `json:"date" bson:"date"`
  Title   string `json:"title" bson:"title"`
  Content string `json:"content" bson:"content"`
}
