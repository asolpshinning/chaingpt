package gpt

import (
	"context"

	"github.com/asolpshinning/db2-warehouse/utils"
	openai "github.com/sashabaranov/go-openai"
)

func ChatGPT(prompt string) (string, error) {
	token := utils.GoDotEnv("token")
	client := openai.NewClient(token)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
