package utils

import (
  "strings"
  "log"

  "gopkg.in/yaml.v2"
  "errors"

  "idleworkshop/website/models"
)

func ParseMarkdown(markdown string) (*models.ParsedBlogPost, error) {
  // Split the markdown into metadata and content
  log.Println(markdown)
  parts := strings.SplitN(markdown, "---", 3)
  log.Println(parts)
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
  contentParts := strings.SplitN(content, "<!-- summary -->", 2)

  summary := strings.TrimSpace(contentParts[0])
  mainContent := ""
  if len(contentParts) > 1 {
    mainContent = strings.TrimSpace(contentParts[1])
  }
  
  return &models.ParsedBlogPost{
    Metadata: metadata,
    Summary: summary,
    Content:  mainContent,
  }, nil
}
