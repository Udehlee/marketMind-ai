package service

import (
	"context"
	"fmt"

	"github.com/Udehlee/marketMind-ai/internals/adapters/llm"
	"github.com/Udehlee/marketMind-ai/internals/core/domain"
)

type RAG struct {
	datasource []domain.DataSource
	llm        *llm.Openai
}

func NewRAG(d []domain.DataSource, llm *llm.Openai) *RAG {
	return &RAG{
		datasource: d,
		llm:        llm,
	}
}

// Generate generates summmary
func (r *RAG) Generate(ctx context.Context) (string, error) {
	var allData []domain.ContentItem

	for _, s := range r.datasource {
		data, err := s.Fetch()
		if err != nil {
			return "", fmt.Errorf("failed to fetch data from source: %w", err)
		}
		allData = append(allData, data...)
	}

	summary, err := r.llm.Summary(allData)
	if err != nil {
		return "", fmt.Errorf("failed to summarize data: %w", err)
	}

	return summary, nil
}
