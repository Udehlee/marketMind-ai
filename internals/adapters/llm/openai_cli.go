package llm

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Udehlee/marketMind-ai/internals/core/domain"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type Openai struct {
	openClient openai.Client
}

func NewOpenai() *Openai {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("Openai api key not set")
	}
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	return &Openai{
		openClient: client,
	}
}

// summary creates a summary text using gpt
func (o *Openai) Summary(contents []domain.ContentItem) (string, error) {
	var builder strings.Builder
	for _, c := range contents {
		builder.WriteString(fmt.Sprintf("Title: %s\nDescription: %s\nSource: %s\n\n",
			c.Title, c.Description, c.Source))
	}

	prompt := fmt.Sprintf(
		"Summarize the following stock and news data into a simple concise summary:\n\n%s",
		builder.String(),
	)

	resp, err := o.openClient.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model: openai.ChatModelGPT4oMini,
	})

	if err != nil {
		return "", fmt.Errorf("failed to call OpenAI: %w", err)
	}

	return resp.Choices[0].Message.Content, nil
}
