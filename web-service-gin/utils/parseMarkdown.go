package utils

import (
  "bytes"
  "strings"

  "gopkg.in/yaml.v2"
  "errors"

  "idleworkshop/website/models"
)

func ParseMarkdown(markdown string) (*models.ParsedBlogPost, error) {
  // Split the markdown into metadata and content
  parts := strings.SplitN(markdown, "---", 3)
  if len(parts) < 3 {
    return nil, errors.New("Invalid markdown format")
  }
  // Parse the metadata
  var metadata models.BlogMetadata
  if err := yaml.Unmarshal([]byte(parts[1]), &metadata); err != nil {
    return nil, err
  }
  // Parse the content
  content := strings.TrimSpace(parts[2])
  return &models.ParsedBlogPost{
    Metadata: metadata,
    Content:  content,
  }, nil
}
