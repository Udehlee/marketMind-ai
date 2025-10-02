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

type MarketstackResponse struct {
	Data []struct {
		Symbol string  `json:"symbol"`
		Close  float64 `json:"close"`
		Volume int64   `json:"volume"`
		Date   string  `json:"date"`
	} `json:"data"`
}
