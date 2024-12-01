package services

import "github.com/nevesgustavo/llm-test/internal/core/ports"

type PromptService struct {
	promptSender ports.LLMPromptSender
}

func NewPromptService(promptClient ports.LLMPromptSender) *PromptService {
	return &PromptService{
		promptSender: promptClient,
	}
}

func (p PromptService) SendMessage(message string) (string, error) {
	return p.promptSender.SendMessage(message)
}
