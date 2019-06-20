package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

// handle the CloudFormation request
func handleRequest(ctx context.Context, inputEvent json.RawMessage) error {

	// Unmarshall here rather than on invocation to ensure we catch errors
	var event CloudFormationCustomResourceEvent
	err := json.Unmarshal(inputEvent, &event)
	if err != nil {
		return fmt.Errorf("Input event error: %v", err)
	}

	// Grab the named parameter value from SSM
	p, err := SSMParameterAccess(event.ResourceProperties.ParameterName)
	if err != nil {
		return fmt.Errorf("SSM access error: %v", err)
	}

	// Build a response struct for CloudFormation
	res := buildCloudFormationResponse(event, p)
	if err != nil {
		return fmt.Errorf("CloudFormation response error: %v", err)
	}

	// Convert the CloudFormation response to JSON
	JSON, err := json.Marshal(res)
	if err != nil {
		return fmt.Errorf("JSON error: %v", err)
	}

	// Create an HTTP client to send the S3 request to the presigned URL
	client := &http.Client{}
	req, err := http.NewRequest("PUT", event.ResponseURL, bytes.NewReader(JSON))
	if err != nil {
		return fmt.Errorf("http request error: %v", err)
	}

	// Send the HTTP request to write out the response data to S3
	req.ContentLength = int64(len(JSON))
	s3res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("S3 request error: %v", err)
	}
	if s3res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("S3 response error: %v", s3res)
	}

	return nil
}

func buildCloudFormationResponse(
	event CloudFormationCustomResourceEvent,
	parameterValue string,
) CloudFormationCustomResourceResponse {
	// Create a response struct
	res := CloudFormationCustomResourceResponse{
		Status:             "SUCCESS",
		NoEcho:             true,
		StackID:            event.StackID,
		RequestID:          event.RequestID,
		LogicalResourceID:  event.LogicalResourceID,
		PhysicalResourceID: event.ResourceProperties.ParameterName,
	}

	// Add the value here as it's a nested struct
	res.Data.ParameterValue = parameterValue

	return res
}

func main() {
	lambda.Start(handleRequest)
}
