![Golang](https://github.com/mathisve/docker-on-lambda/actions/workflows/go.yaml/badge.svg)
![Docker](https://github.com/mathisve/docker-on-lambda/actions/workflows/docker.yaml//badge.svg)

# docker-on-lambda
This repository contains the project files of [this video](https://youtu.be/EYqFbRsh_RM).
Read [this article](https://docs.aws.amazon.com/lambda/latest/dg/go-image.html) for a more in depth text-based tutorial.

## commands
Golang
```
# module name cannot be time (because that package already exits!)
go mod init lambda
go mod tidy
```

Ecr
```
# optionally you can create the container repository via the AWS CLI
aws ecr create-repository --repository-name time
```


Base docker image building and pushing
```
docker build -t time:latest .

# use your own account id obviously
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 123456789012.dkr.ecr.us-east-1.amazonaws.com

docker tag  time:latest 123456789012.dkr.ecr.us-east-1.amazonaws.com/time:latest

docker push 123456789012.dkr.ecr.us-east-1.amazonaws.com/time:latest
```

RIE (Runtime Interface Emulator)
```
mkdir -p ~/.aws-lambda-rie && curl -Lo ~/.aws-lambda-rie/aws-lambda-rie \
https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie \
&& chmod +x ~/.aws-lambda-rie/aws-lambda-rie

docker run -d -v ~/.aws-lambda-rie:/aws-lambda --entrypoint /aws-lambda/aws-lambda-rie  -p 9000:8080 time:latest /main

curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{}'
```
