package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/nevesgustavo/llm-test/internal/core/ports"
	"github.com/nevesgustavo/llm-test/internal/dto"
)

type PromptHandler struct {
	service ports.LLMPromptSender
}

func NewLLMHandler(service ports.LLMPromptSender) *PromptHandler {
	return &PromptHandler{
		service: service,
	}
}

func (m PromptHandler) HandlePrompt(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req, err := decodeBody(request)
	if err != nil {
		return createJSONResponse(
			http.StatusBadRequest,
			"Not a valid request",
		)
	}

	startTime := time.Now()

	resp, err := m.service.SendMessage(req.Prompt)
	if err != nil {
		return createDefaultJSONResponse(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	timeElapsed := time.Since(startTime)

	response := dto.PromptResponse{
		Response:    resp,
		ElapsedTime: timeElapsed.String(),
	}

	return createJSONResponse(
		http.StatusOK,
		response,
	)
}

func decodeBody(request events.APIGatewayProxyRequest) (*dto.Prompt, error) {
	var req *dto.Prompt
	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return nil, err
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}
	return req, nil
}

func createJSONResponse(
	statusCode int,
	body interface{},
) (events.APIGatewayProxyResponse, error) {
	binaryBody, err := json.Marshal(body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	slog.Info("prompt response:", map[string]any{
		"status_code": statusCode,
		"body":        string(binaryBody),
	})

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(binaryBody),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func createDefaultJSONResponse(
	statusCode int,
	body any,
) (events.APIGatewayProxyResponse, error) {
	binaryBody, err := json.Marshal(body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	slog.Info("prompt sender response:", map[string]any{
		"status_code": statusCode,
		"body":        string(binaryBody),
	})

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(binaryBody),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
