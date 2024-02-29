# AWS Lambda + API Gateway + Goで作る簡単なサーバレス構成

## Install Serverless Framework

```nash
npm i -g serverless
```

## Build

```bash
GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/go-lambda-api main.go
```

## Run

```bash
go run main.go
```

## Deploy

```bash
sls deploy
```

## Delete

```bash
sls remove
```