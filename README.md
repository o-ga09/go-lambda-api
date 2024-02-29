# AWS Lambda + API Gateway + Goで作る簡単なサーバレス構成

## Build

```bash
GOOS=linux go build -ldflags="-s -w" -o bin/go-lambda-api main.go
```

## Run

```bash
go run main.go
```

## Deploy

```bash
sls deploy
```