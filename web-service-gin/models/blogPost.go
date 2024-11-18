type BlogPost struct {
  ID      int    `json:"id"`
  Title   string `json:"title"`
  Content string `json:"content"`
  Tags  []Tag  `json:"tags"`
  Date    string `json:"date"`
}
