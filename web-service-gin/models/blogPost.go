package models

import (
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPost struct {
  ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
  Date    primitive.DateTime `json:"date" bson:"date"`
  Title   string `json:"title" bson:"title"`
  Slug    string `json:"slug" bson:"slug"`
  Tags    []string `json:"tags" bson:"tags"`
  Nutshell string `json:"nutshell" bson:"nutshell"`
  ReadTime string `json:"readTime" bson:"readTime"`
  Summary string `json:"summary" bson:"summary"`
  Content string `json:"content" bson:"content"`
}

type BlogMetadata struct {
  Title string `yaml:"title"`
  Tags []string `yaml:"tags"`
  Nutshell string `yaml:"nutshell"`
  Topic string `yaml:"topic"`
  ReadTime string `yaml:"readTime"`
}

type ParsedBlogPost struct {
  Metadata BlogMetadata
  Summary string
  Content string
}

