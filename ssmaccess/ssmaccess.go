package ssmaccess

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// SSMParameterAccess provides access to SSM parameters
func SSMParameterAccess(parameterName string) (string, error) {

	sess := session.Must(session.NewSession())
	svc := ssm.New(sess)

	paramInput := &ssm.GetParameterInput{
		Name:           aws.String(parameterName),
		WithDecryption: aws.Bool(true),
	}

	p, err := svc.GetParameter(paramInput)
	if err != nil {
		return "", err
	}

	return *p.Parameter.Value, nil
}
