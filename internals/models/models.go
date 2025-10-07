package models

type Feed struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

type FeedItem struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Content     string `json:"content"`
	PublishedAt string `json:"published"`
}

type FeedResult struct {
	Feed  Feed
	Items []FeedItem
}

type FeedConfig struct {
	Feeds []Feed `yaml:"feeds"`
}

type AlphaVantageResponse struct {
	MetaData        map[string]string    `json:"Meta Data"`
	TimeSeriesDaily map[string]DailyData `json:"Time Series (Daily)"`
}

type DailyData struct {
	Open   float64 `json:"1. open,string"`
	High   float64 `json:"2. high,string"`
	Low    float64 `json:"3. low,string"`
	Close  float64 `json:"4. close,string"`
	Volume int64   `json:"5. volume,string"`
}
