package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.CognitoEventUserPoolsPreSignup) (events.CognitoEventUserPoolsPreSignup, error) {
	// Log the received event
	fmt.Printf("Received event: %+v\n", event)

	// Auto-verify the email and confirm the user
	event.Response.AutoVerifyEmail = true
	event.Response.AutoConfirmUser = true

	return event, nil
}

func main() {
	lambda.Start(handler)
}
