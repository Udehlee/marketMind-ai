package domain

import "time"

type ContentItem struct {
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Source      string                 `json:"source"`
	Timestamp   time.Time              `json:"timestamp"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type DataSource interface {
	Fetch() ([]ContentItem, error)
}
