package main

import (
	"context"
	"rando/characters"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type Seed struct {
	Seed int64 `json:"seed"`
}

func HandleRequest(ctx context.Context, seed Seed) ([]byte, error) {
	var s int64
	if seed.Seed == 0 {
		s = time.Now().UnixNano()
	} else {
		s = seed.Seed
	}
	return characters.GenerateCharacter(s), nil
}

func main() {
	lambda.Start(HandleRequest)
}
