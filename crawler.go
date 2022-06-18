package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func Crawl() (string, error) {
	env := os.Getenv("ENV")
	if env != "prod" {
		log.Println("Mocking request to USCIS...")
		return `
		<html>
			<body>
				<div class="rows text-center">
					  <h1>Case Is Being Actively Reviewed By USCIS</h1>
					  <p>Placeholder</p>
				</div>
			</body>
		</html>
		`, nil
	}
	sourceURL := viper.Get("USCIS_URL").(string)
	receiptNumber := viper.Get("RECEIPT_NUMBER").(string)
	formData := url.Values{
		"upcomingActionsCurrentPage":  {"0"},
		"completedActionsCurrentPage": {"0"},
		"appReceiptNum":               {receiptNumber},
		"caseStatusSearchBtn":         {"CHECK STATUS"},
		"changeLocale":                {},
	}
	client := &http.Client{}

	req, err := http.NewRequest("POST", sourceURL,
		strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Fatalf("Request error %v\n", err)
		return "", err
	}

	log.Println("Making real request to USCIS...")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Resposne error %v\n", err)
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Response body parsing error %v\n", err)
		return "", err
	}

	return string(body), nil
}
