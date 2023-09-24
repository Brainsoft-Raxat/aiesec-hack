package repository

import (
	"context"
	"fmt"

	"github.com/Brainsoft-Raxat/aiesec-hack/internal/app/config"
	"github.com/Brainsoft-Raxat/aiesec-hack/internal/models"
	"github.com/sashabaranov/go-openai"
)

type gptRepo struct {
	client *openai.Client
}

func NewGPTRepository(cfg config.GPT) GPT {
	return &gptRepo{
		client: openai.NewClient(cfg.Token),
	}
}

func (r *gptRepo) SendPrompt(ctx context.Context, data []models.Event) (string, error) {
	var eventData string
	for i, event := range data {
		eventData += fmt.Sprintf("%d. %s - %s\n   Description: %s\n   Location: %s\n   Date and Time: %s\n\n", i+1, event.Title, event.Author, event.Description, event.Address, event.Datetime)
	}

	userMessage := fmt.Sprintf("I'm looking for events to attend today. Can you suggest something?\n\n%s", eventData)

	resp, err := r.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a helpful assistant that can suggest upcoming events for today.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userMessage,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
