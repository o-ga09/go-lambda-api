package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := Response{
		Count:  1,
		Limit:  1,
		Offset: 10,
		User: User{
			Name: "test",
			Age:  "29",
		},
	}
	json, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{
		Body:       string(json),
		StatusCode: 200,
	}, nil
}

type Response struct {
	Count  int  `json:"count,omitempty"`
	Limit  int  `json:"limit,omitempty"`
	Offset int  `json:"offset,omitempty"`
	User   User `json:"user,omitempty"`
}

type User struct {
	Name string
	Age  string
}
