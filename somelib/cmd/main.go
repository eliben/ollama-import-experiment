package main

import (
	"context"
	"fmt"
	"log"

	"github.com/eliben/ollama-import-experiment/somelib"
	"github.com/jmorganca/ollama/api"
)

// This should work just fine - there's only one "api" package here, from
// jmorganca/ollama/api - it's used by MakeRequest and here.
func main() {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	req := somelib.MakeRequest("gemma", "what is the closest star to the sun?")

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
