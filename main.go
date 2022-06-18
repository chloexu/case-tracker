package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/spf13/viper"
)

type Event struct {
	Name string `json:"name"`
}

func lambdaHandler(ctx context.Context, name Event) (string, error) {
	return startJob()
}

func startJob() (string, error) {
	env := os.Getenv("ENV")
	log.Printf("Loading config file for %v", env)
	switch {
	case env == "prod":
		viper.SetConfigName("prod")
	default:
		viper.SetConfigName("default")
	}
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
		return "", err
	}

	pageSource, err := Crawl()
	if err != nil {
		return "", err
	}
	result, err := Parse(pageSource)
	if err != nil {
		return "", err
	}
	ok, err := Message(result)
	if err != nil {
		return "", err
	}
	return ok, nil
}

func main() {
	env := os.Getenv("ENV")
	if env == "prod" {
		// running in lambda mode
		lambda.Start(lambdaHandler)
	} else {
		// running in local without AWS context
		startJob()
	}
}
