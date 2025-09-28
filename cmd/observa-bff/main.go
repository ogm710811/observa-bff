package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Message string `json:"message"`
}

type Response struct {
	Greeting string `json:"greeting"`
}

func handler(ctx context.Context, event Event) (Response, error) {
	// Here, you process the request. "event" contains input from AWS
	greeting := "Hello and welcome, " + event.Message + "!"
	return Response{Greeting: greeting}, nil
}

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	lambda.Start(handler)
}
