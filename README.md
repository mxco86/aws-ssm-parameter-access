# AWS SSM Parameter Access

## What Does This Do?
A CloudFormation custom resource backed by a Lambda function that allows
access to an encrypted SSM parameter by the parameter name

## Usage
Once deployed to a lambda function the custom resource can be used in
CloudFormation templates to access SSM parameters.

```YAML
---
AWSTemplateFormatVersion: "2010-09-09"
Description: "Some CloudFormation-Defined System"

Resources:
  # Define the Custom Resource to access SSM parameters
  SecretParameter:
    Type: AWS::CloudFormation::CustomResource
    Properties:
      ServiceToken: !Sub arn:aws:lambda:${AWS::Region}:${AWS::AccountId}:function:SSMParameterAccess
      ParameterName: "MySecretParameterName"

  # Use the parameter value in a resource definition
  AnotherResource:
    Secret: !GetAtt MySecretParameterName.ParameterValue

```

## Configuration
There is no explicit configuration of the code. The custom resource provides
access to any SSM parameter store parameters that are present in the AWS
account the resource is deployed to. The IAM policy in the SAM configuration
allows ssm:GetParameter access. The resource must also have access to the key
used to encrypt parameters. If the default AWS key is used access is in place
by default. If another key is used IAM permission must be given to use this
key.

## Build and Deployment
The Go binary must be built before deployment with the correct
compilation flags for Lambda execution

```sh
# Build Go binary for Lambda execution
GOOS=linux go build -ldflags="-s -w"
```

The Lambda stack is defined in a SAM configuration and can be built and
deployed using the standard SAM commands. The S3 bucket that holds the
deployment package can be any existing bucket.

```sh
# Stack creation
sam package --template-file sam.yaml --s3-bucket ${DeploymentBucketName} --output-template-file sam-pkg.yaml
sam deploy --template-file ./sam-pkg.yaml --stack-name ${StackName} --capabilities CAPABILITY_IAM

# Stack deletion
aws cloudformation delete-stack --stack-name ${StackName}
```

## Testing
The code can be tested locally using the standard SAM commands or directly
using a lambci docker container. A presigned S3 URL is needed to receive the
response

```sh
docker run --rm -v "$PWD":/var/task \
    -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY \
    -e AWS_SESSION_TOKEN -e AWS_SECURITY_TOKEN \
    -e AWS_REGION=eu-west-1 \
    lambci/lambda:go1.x aws-ssm-parameter-access \
    "{
        \"RequestType\" : \"Create\",
        \"ResponseURL\" : \"https://presignedS3URL/\",
        \"StackId\" : \"arn:aws:cloudformation:us-west-2:123456789012:stack/stack-name/guid\",
        \"RequestId\" : \"unique id for this create request\",
        \"ResourceType\" : \"Custom::TestResource\",
        \"LogicalResourceId\" : \"MyTestResource\",
        \"ResourceProperties\" : {
            \"ParameterName\" : \"ATestParameter\"
        }
    }"

```
