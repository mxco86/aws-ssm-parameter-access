package awstypes

// CloudFormationCustomResourceEvent -- an event
type CloudFormationCustomResourceEvent struct {
	RequestType        string `json:"RequestType"`
	ResponseURL        string `json:"ResponseURL"`
	StackID            string `json:"StackId"`
	RequestID          string `json:"RequestId"`
	ResourceType       string `json:"ResourceType"`
	LogicalResourceID  string `json:"LogicalResourceId"`
	ResourceProperties struct {
		ParameterName string `json:"ParameterName"`
	} `json:"ResourceProperties"`
}

// CloudFormationCustomResourceResponse -- a response
type CloudFormationCustomResourceResponse struct {
	Status             string `json:"Status"`
	NoEcho             bool   `json:"NoEcho"`
	StackID            string `json:"StackId"`
	RequestID          string `json:"RequestId"`
	LogicalResourceID  string `json:"LogicalResourceId"`
	PhysicalResourceID string `json:"PhysicalResourceId"`
	Data               struct {
		ParameterValue string `json:"ParameterValue"`
	} `json:"Data"`
}
