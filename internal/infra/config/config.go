package config

import "os"

const defaultPromptModel = "gpt-3.5-turbo"

var configuration *Config

type Config struct {
	Model string `json:"model"`
}

func init() {
	err := LoadConfig()
	if err != nil {
		panic(err)
	}
}

func LoadConfig() error {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("the environment variable OPENAI_API_KEY is required")
	}

	config := &Config{
		Model: os.Getenv("PROMPT_MODEL"),
	}

	if config.Model == "" {
		config.Model = defaultPromptModel
	}

	setConfig(config)

	return nil
}

func setConfig(configs *Config) {
	configuration = configs
}

func GetConfig() *Config {
	return configuration
}
