package main

import (
	"context"
	"fmt"
	"log"

	"github.com/eliben/ollama-import-experiment/somelib"
	"github.com/ollama/ollama/api"
)

// This will not work because somelib.MakeRequest returns
// *github.com/jmorganca/ollama/api.GenerateRequest, which
// is a different type for the Go compiler from
// *github.com/ollama/ollama/api.GenerateResponse
func main() {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	req := somelib.MakeRequest("gemma", "how many planets are there?")

	ctx := context.Background()
	respFunc := func(resp api.GenerateResponse) error {
		// Only print the response here; GenerateResponse has a number of other
		// interesting fields you want to examine.
		fmt.Println(resp.Response)
		return nil
	}

	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}
}
