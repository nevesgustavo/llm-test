package prompt

import (
	"context"

	"github.com/nevesgustavo/llm-test/internal/infra/config"
	"github.com/tmc/langchaingo/llms/openai"
)

type PromptSender struct{}

func NewPromptSender() *PromptSender {
	return &PromptSender{}
}

func (p PromptSender) SendMessage(message string) (string, error) {
	return p.send(message)
}

func (PromptSender) send(prompt string) (string, error) {
	llm, err := openai.New(openai.WithModel(config.GetConfig().Model))

	completion, err := llm.Call(context.Background(), prompt)
	if err != nil {
		return "", err
	}

	return completion, nil
}
