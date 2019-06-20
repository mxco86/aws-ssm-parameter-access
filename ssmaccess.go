package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// SSMParameterAccess provides access to a single named SSM parameter
func SSMParameterAccess(parameterName string) (string, error) {

	// Create a new AWS session
	sess := session.Must(session.NewSession())
	svc := ssm.New(sess)

	// Grab the named encrypted parameter
	paramInput := &ssm.GetParameterInput{
		Name:           aws.String(parameterName),
		WithDecryption: aws.Bool(true),
	}

	p, err := svc.GetParameter(paramInput)
	if err != nil {
		return "", err
	}

	// Return just the parameter value
	return *p.Parameter.Value, nil
}
