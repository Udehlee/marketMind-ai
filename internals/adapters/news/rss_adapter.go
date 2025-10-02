package news

import (
	"github.com/Udehlee/marketMind-ai/internals/core/domain"
)

type rssAdapter struct{}

func NewRSSAdapter() *rssAdapter {
	return &rssAdapter{}
}

func (r *rssAdapter) Fetch() ([]domain.ContentItem, error) {
	results, err := FetchAllFeeds()
	if err != nil {
		return nil, err
	}

	var items []domain.ContentItem
	for _, fr := range results {
		items = append(items, rssToContentItems(fr)...)
	}
	return items, nil
}
