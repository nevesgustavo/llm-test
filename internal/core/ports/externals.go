package ports

type LLMPromptSender interface {
	SendMessage(message string) (string, error)
}
