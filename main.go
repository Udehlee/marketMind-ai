package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Udehlee/marketMind-ai/internals/adapters/llm"
	"github.com/Udehlee/marketMind-ai/internals/adapters/market"
	"github.com/Udehlee/marketMind-ai/internals/adapters/news"
	"github.com/Udehlee/marketMind-ai/internals/core/domain"
	"github.com/Udehlee/marketMind-ai/internals/core/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("find out about stock news today: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	news := news.NewRssAdapter()
	market := market.NewMarketAdapter(input)
	llm := llm.NewOpenai()

	rag := service.NewRAG([]domain.DataSource{news, market}, llm)

	answer, err := rag.Generate(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\n", answer)

}
