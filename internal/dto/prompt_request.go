package dto

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Prompt struct {
	Prompt string `json:"prompt"`
}

func (p Prompt) Validate() error {
	errors := map[string]string{}

	if len(strings.Trim(p.Prompt, " ")) == 0 {
		errors["message"] = "prompt is required"
	}

	if len(errors) > 0 {
		bs, _ := json.Marshal(errors)
		return fmt.Errorf("validation errors: %v", string(bs))
	}

	return nil
}
