package main

import (
	"log"
	"strings"

	"golang.org/x/net/html"
)

func Parse(source string) string {
	var result string
	doc, err := html.Parse(strings.NewReader(source))
	if err != nil {
		log.Fatalln("Error parsing html source")
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "h1" {
			result = n.FirstChild.Data
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	log.Printf("Latest status is: %v", result)
	return result
}
