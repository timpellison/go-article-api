package dto

type Article struct {
	Metadata    []Hypermedia `json:"_metadata,omitempty"`
	ID          uint         `json:"id,omitempty"`
	Title       string       `json:"title,omitempty"`
	Description string       `json:"description,omitempty"`
	Content     string       `json:"content,omitempty"`
}

func NewArticle(id uint, title, description, content string, metadata []Hypermedia) *Article {
	return &Article{metadata, id, title, description, content}
}

type ArticleData struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Content     string `json:"content,omitempty"`
}
