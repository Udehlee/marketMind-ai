package market

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Udehlee/marketMind-ai/internals/core/domain"
	"github.com/Udehlee/marketMind-ai/internals/models"
)

type marketAdapter struct {
	apiKey  string
	symbols []string
}

func NewMarketAdapter(symbols []string) *marketAdapter {
	apiKey := os.Getenv("ALPHA_VANTAGE_API_KEY")
	if apiKey == "" {
		panic("Alpha Api key not set")
	}

	return &marketAdapter{
		apiKey:  apiKey,
		symbols: symbols,
	}
}

// Fetch gets stock data from alphavantage
// normalizes it to ContentItem
func (m *marketAdapter) Fetch() ([]domain.ContentItem, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s",
		m.symbols, m.apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch market data: %w", err)
	}
	defer resp.Body.Close()

	var alphaRes models.AlphaVantageResponse
	if err := json.NewDecoder(resp.Body).Decode(&alphaRes); err != nil {
		return nil, fmt.Errorf("failed to decode market data: %w", err)
	}

	var items []domain.ContentItem
	for date, data := range alphaRes.TimeSeriesDaily {
		items = append(items, domain.ContentItem{
			Title:       fmt.Sprintf("%s %s: %.2f", m.symbols, date, data.Close),
			Description: fmt.Sprintf("Stock %s closed at %.2f (open %.2f, high %.2f, low %.2f)", m.symbols, data.Close, data.Open, data.High, data.Low),
			Source:      "alphavantage",
			Timestamp:   time.Now(),
			Metadata: map[string]interface{}{
				"symbol": m.symbols,
				"open":   data.Open,
				"close":  data.Close,
				"high":   data.High,
				"low":    data.Low,
				"volume": data.Volume,
				"date":   date,
			},
		})
		break
	}

	return items, nil
}
