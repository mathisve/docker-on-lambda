package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context) (string, error) {
	t := time.Now().String()

	return t, nil
}