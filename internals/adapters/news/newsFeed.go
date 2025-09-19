package news

import (
	"net/http"
	"os"
	"time"

	"github.com/Udehlee/marketMind-ai/internals/models"
	"github.com/mmcdole/gofeed"
	"gopkg.in/yaml.v2"
)

func parseFeeds(path string) (*models.FeedConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var FeedCfg models.FeedConfig
	if err := yaml.Unmarshal(data, &FeedCfg); err != nil {
		return nil, err
	}
	return &FeedCfg, nil
}

func fetchFeed(feed models.Feed) (models.FeedResult, error) {
	var (
		client = &http.Client{Timeout: 10 * time.Second}
		parser = gofeed.NewParser()
	)

	parser.Client = client
	parsed, err := parser.ParseURL(feed.Url)
	if err != nil {
		return models.FeedResult{}, err
	}

	var Item []models.FeedItem
	for _, i := range parsed.Items {
		Item = append(Item, models.FeedItem{
			Title:       i.Title,
			Link:        i.Link,
			Content:     i.Content,
			PublishedAt: i.Published,
		})

	}

	res := models.FeedResult{
		Feed:  feed,
		Items: Item,
	}

	return res, nil
}

func FetchAllFeeds() ([]models.FeedResult, error) {
	cfg, err := parseFeeds("config/feeds.yaml")
	if err != nil {
		return nil, err
	}

	resCh := make(chan models.FeedResult)
	fetchErr := make(chan error)

	for _, feed := range cfg.Feeds {
		go func(f models.Feed) {
			result, err := fetchFeed(f)
			if err != nil {
				fetchErr <- err
				return
			}
			resCh <- result
		}(feed)
	}

	var results []models.FeedResult
	for range cfg.Feeds {
		select {
		case res := <-resCh:
			results = append(results, res)
		case err := <-fetchErr:
			return nil, err
		}
	}

	return results, nil
}
