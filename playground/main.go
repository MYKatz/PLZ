package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/MYKatz/PLZ/interpreter"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Code string `json:"code"`
}

type Response struct {
	Res        string `json:"res"`
	StatusCode int    `json:"statusCode"`
}

func interpret(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	resp := events.APIGatewayProxyResponse{Headers: make(map[string]string)}
	resp.Headers["Access-Control-Allow-Origin"] = "*"
	//capture output of stdout
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout = w

	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	code := req.QueryStringParameters["code"]
	interpreter.Interpret(code)

	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	headers := make(map[string]string)
	headers["Access-Control-Allow-Origin"] = "*"

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    headers,
		Body:       string(out),
	}, nil
}

func main() {
	lambda.Start(interpret)
}
