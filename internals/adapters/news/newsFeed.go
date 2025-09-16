package news

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

type FeedConfig struct {
	Feed  Feed
	Items []FeedItem
}
