package ai

import (
	"context"
	"fmt"
	"github.com/notopos/api/internal/config"
	"github.com/sashabaranov/go-openai"
)

type AIService interface {
	GenerateSummary(ctx context.Context, data string) (string, error)
	GetRecommendation(ctx context.Context, userID string, context string) (string, error)
}

type aiService struct {
	client *openai.Client
	config *config.Config
}

func NewAIService(cfg *config.Config) AIService {
	client := openai.NewClient(cfg.OpenAIKey)
	return &aiService{
		client: client,
		config: cfg,
	}
}

func (s *aiService) GenerateSummary(ctx context.Context, data string) (string, error) {
	resp, err := s.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a business analyst assistant for NOTOPOS AI. Provide a concise summary of the sales data provided.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: data,
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("AI error: %v", err)
	}

	return resp.Choices[0].Message.Content, nil
}

func (s *aiService) GetRecommendation(ctx context.Context, userID string, contextStr string) (string, error) {
	// Logic for RAG or simple prompt based recommendation
	return "Sample recommendation based on context", nil
}
