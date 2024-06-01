package main

import (
	"context"
	"datalogger/modal"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func update(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var book modal.Books
	err := json.Unmarshal([]byte(req.Body), &book)
	fmt.Println("book", book)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       err.Error(),
		}, nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while retrieving AWS credentials",
		}, nil
	}

	svc := dynamodb.NewFromConfig(cfg)
	newBook, err := svc.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String("books"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: req.PathParameters["id"],
			},
		},
		ExpressionAttributeNames: map[string]string{
			"#name":   "name",
			"#author": "author",
		},
		UpdateExpression: aws.String("SET #name = :n, #author = :a"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":n": &types.AttributeValueMemberS{
				Value: book.Name,
			},
			":a": &types.AttributeValueMemberS{
				Value: book.Author,
			},
		},
		ReturnValues: types.ReturnValueAllNew,
	})

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	res, _ := json.Marshal(newBook)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(res),
	}, nil
}

func main() {
	lambda.Start(update)
}
