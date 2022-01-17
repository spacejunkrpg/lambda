package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/spacejunkrpg/character"
)

type Seed struct {
	Seed int64 `json:"seed"`
}

type Response struct {
	StatusCode int    `json:"statusCode"`
	Base64     bool   `json:"isBase64Encoded"`
	Body       string `json:"body"`
}

func HandleRequest(ctx context.Context, seed Seed) (Response, error) {
	var s int64
	if seed.Seed == 0 {
		s = time.Now().UnixNano()
	} else {
		s = seed.Seed
	}

	r := Response{}
	c := character.GenerateCharacter(s)
	j, _ := json.Marshal(c)
	r.Body = string(j)
	r.Base64 = false
	r.StatusCode = 200

	return r, nil
}

func main() {
	lambda.Start(HandleRequest)
}
