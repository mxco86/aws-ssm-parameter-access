* AWS-SSM-Parameter-Access

```sh

# Run locally
docker run --rm -v "$PWD":/var/task \
    -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY \
    -e AWS_SESSION_TOKEN -e AWS_SECURITY_TOKEN \
    -e AWS_REGION=eu-west-1 \
    lambci/lambda:go1.x aws-ssm-access "{ \"parameterName\": \"AParameterName\" }"

```
