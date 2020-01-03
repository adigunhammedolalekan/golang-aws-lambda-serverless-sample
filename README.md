## Golang AWS Lambda example
This is an example of how to create REST API using AWS Lambda and serverless Framework

### Requirements
* Go 1.1x
* aws-cli
* serverless framework

### How to run

* Configure AWS credentials by running `aws configure` and follow the prompt
* Run `make deploy` - This will build functions binary and use serverless to deploy it