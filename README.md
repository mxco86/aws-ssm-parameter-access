* AWS-SSM-Parameter-Access

```sh

# Run locally
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
