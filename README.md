![Golang](https://github.com/mathisve/docker-on-lambda/actions/workflows/go.yaml/badge.svg)
![Docker](https://github.com/mathisve/docker-on-lambda/actions/workflows/docker.yaml//badge.svg)

# docker-on-lambda
This repository contains the project files of [this video](https://youtu.be/EYqFbRsh_RM).
Read [this article](https://docs.aws.amazon.com/lambda/latest/dg/go-image.html) for a more in depth text-based tutorial.

## commands
### Golang
```bash
go mod init lambda-function
go mod tidy
```

### ECR
optionally you can create the container repository via the AWS Console
```bash
aws ecr create-repository --repository-name lambda-function
```


Base docker image building and pushing
```bash
docker build -t lambda-function:latest .

aws ecr get-login-password --region {region} | docker login --username AWS --password-stdin {account-number}.dkr.ecr.{region}.amazonaws.com

docker tag lambda-function:latest {account-number}.dkr.ecr.{region}.amazonaws.com/lambda-function:latest

docker push {account-number}.dkr.ecr.{region}.amazonaws.com/lambda-function:latest
```

RIE (Runtime Interface Emulator)
```bash
mkdir -p ~/.aws-lambda-rie && curl -Lo ~/.aws-lambda-rie/aws-lambda-rie \
https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie \
&& chmod +x ~/.aws-lambda-rie/aws-lambda-rie

docker run -d -v ~/.aws-lambda-rie:/aws-lambda --entrypoint /aws-lambda/aws-lambda-rie  -p 9000:8080 lambda-function:latest /main

curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{}'
```
