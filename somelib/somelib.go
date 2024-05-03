package somelib

import "github.com/jmorganca/ollama/api"

func MakeRequest(model string, prompt string) *api.GenerateRequest {
	return &api.GenerateRequest{
		Model:  model,
		Prompt: prompt,

		Stream: new(bool),
	}
}
