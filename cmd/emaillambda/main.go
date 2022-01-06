package main

import (
	"bytes"
	"context"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"github.com/4thel00z/libemail/pkg/v1/libemail"
	"github.com/4thel00z/libemail/pkg/v1/libemail/gmail"
	"github.com/4thel00z/libemail/pkg/v1/libemail/senders"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	g         senders.GmailSender
	basicAuth string
)

func init() {
	basicAuth = os.Getenv("BASIC_AUTH")

	g = senders.GmailSender{}
	token, err := gmail.LoadTokenFromEnv("GMAIL_TOKEN")
	if err != nil {
		log.Fatalln(err.Error())
	}
	config, err := gmail.GoogleConfigFromEnvVar("GMAIL_CONFIG")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = g.Init(config, token)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	auth := request.Headers["authorization"]
	if 1 != subtle.ConstantTimeCompare([]byte(auth), []byte(basicAuth)) {
		log.Println("received: ", auth)
		log.Println("expected: ", basicAuth)
		return events.APIGatewayProxyResponse{
			Body: "",
			Headers: map[string]string{
				"WWW-Authenticate": "Basic",
			},
			StatusCode: http.StatusUnauthorized}, errors.New("not authorized")

	}
	var email *libemail.Email
	err := json.NewDecoder(strings.NewReader(request.Body)).Decode(email)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "", StatusCode: http.StatusBadRequest}, err
	}

	res, err := g.Send(email)
	log.Printf("Response:\n%#v", res)
	
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}
	buffer := bytes.NewBuffer([]byte{})
	err = json.NewEncoder(buffer).Encode(res)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       buffer.String(),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
