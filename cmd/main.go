package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nevesgustavo/llm-test/internal/core/services"
	"github.com/nevesgustavo/llm-test/internal/externals/prompt"
	"github.com/nevesgustavo/llm-test/internal/handlers"
)

func main() {
	handler := handlers.NewLLMHandler(
		services.NewPromptService(
			prompt.NewPromptSender(),
		),
	)
	lambda.Start(handler.HandlePrompt)
}
