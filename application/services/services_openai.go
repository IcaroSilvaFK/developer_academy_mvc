package services

import (
	"context"
	"fmt"
	"os"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	openai "github.com/sashabaranov/go-openai"
)

type AIService struct {
	context context.Context
	c       *openai.Client
}

type AIServiceInterface interface {
	VerifyIfIsValidChallenge(title string) bool
	GenerateHintFromChallenge(title string) (string, error)
}

func NewAIService() AIServiceInterface {

	client := openai.NewClient(os.Getenv(utils.OPEN_AI_API_KEY))

	return &AIService{
		context.Background(), client,
	}
}

func (as *AIService) VerifyIfIsValidChallenge(title string) bool {

	res, err := as.c.Moderations(
		as.context,
		openai.ModerationRequest{
			Model: openai.ModerationTextStable,
			Input: title,
		},
	)

	if err != nil {
		return false
	}

	for _, r := range res.Results {
		if r.Categories.Sexual || r.Categories.Violence || r.Categories.Hate {
			return false
		}
	}

	return true
}

func (as *AIService) GenerateHintFromChallenge(title string) (string, error) {

	res, err := as.c.CreateChatCompletion(
		as.context,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("De uma lista com 3 items  sobre '%s' somente a lista", title),
				},
			},
		},
	)

	if err != nil {
		return "", nil
	}

	return res.Choices[0].Message.Content, nil
}
