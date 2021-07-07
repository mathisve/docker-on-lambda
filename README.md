# docker-on-lambda
This repository meant to be "enjoyed" alongside [this]() video (which hasn't been released yet.)

## commands
Golang
```
# cannot be time (because that package already exits!)
go mod init lambda
go mod tidy
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