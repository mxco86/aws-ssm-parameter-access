package main

import (
	"./ssmaccess"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

type event struct {
	ParameterName string `json:"parameterName"`
}

func handleRequest(ctx context.Context, event event) (string, error) {
	p, err := ssmaccess.SSMParameterAccess(event.ParameterName)
	if err != nil {
		return "", err
	}

	return p, nil
}

func main() {
	lambda.Start(handleRequest)
}
