package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func main() {
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
	}

	pageSource := Crawl()
	result := Parse(pageSource)
	Message(result)
}
