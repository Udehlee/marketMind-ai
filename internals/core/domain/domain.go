package internals

import "time"

type MarketInfo struct {
	Symbol    string
	Price     float32
	Volume    int64
	Timestamp time.Time
}

type NewsInfo struct {
	Headline  string
	Source    string
	Content   string
	Timestamp time.Time
}

type Insight struct {
	Symbol    string
	Message   string
	Timestamp time.Time
}

type MarketFeed interface {
	WatchMarket(symbol string) (MarketInfo, error)
}

type NewsFeed interface {
	FetchNews() ([]NewsInfo, error)
}
