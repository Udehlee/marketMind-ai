package market

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Udehlee/marketMind-ai/internals/core/domain"
	"github.com/Udehlee/marketMind-ai/internals/models"
)

type marketAdapter struct {
	apiKey  string
	symbols []string
}

func NewMarketAdapter(apiKey string, symbols []string) *marketAdapter {
	return &marketAdapter{
		apiKey:  apiKey,
		symbols: symbols,
	}
}

// Fetch gets stock data from Marketstack
// normalizes it to ContentItem
func (m *marketAdapter) Fetch() ([]domain.ContentItem, error) {
	url := fmt.Sprintf("http://api.marketstack.com/v1/eod?access_key=%s&symbols=%s",
		m.apiKey, m.symbols)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch market data: %w", err)
	}
	defer resp.Body.Close()

	var MarketRes models.MarketstackResponse
	if err := json.NewDecoder(resp.Body).Decode(&MarketRes); err != nil {
		return nil, fmt.Errorf("failed to decode market data: %w", err)
	}

	var items []domain.ContentItem
	for _, d := range MarketRes.Data {
		items = append(items, domain.ContentItem{
			Title:       fmt.Sprintf("%s closing price: %.2f", d.Symbol, d.Close),
			Description: fmt.Sprintf("Stock %s closed at %.2f with volume %d", d.Symbol, d.Close, d.Volume),
			Source:      "marketstack",
			Timestamp:   time.Now(),
			Metadata: map[string]interface{}{
				"symbol": d.Symbol,
				"close":  d.Close,
				"volume": d.Volume,
				"date":   d.Date,
			},
		})
	}

	return items, nil
}
