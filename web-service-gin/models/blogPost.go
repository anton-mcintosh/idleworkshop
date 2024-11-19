package models

type BlogPost struct {
  ID      int    `json:"id" bson:"id"`
  Date    string `json:"date" bson:"date"`
  Title   string `json:"title" bson:"title"`
  Content string `json:"content" bson:"content"`
}
