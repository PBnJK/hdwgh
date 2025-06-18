package main

import (
	"flag"
	"log"

	"github.com/pbnjk/hdwgh/internal/article"

	"github.com/pbnjk/hdwgh/pkg/backtracker"
)

func main() {
	flagInput := flag.String("input", "", "Initial article")

	flag.Parse()

	initialArticle, err := article.New(*flagInput)
	if err != nil {
		log.Fatalf("An error occurred getting the initial article: %v", err)
	}

	bt := backtracker.New(initialArticle, backtracker.DefaultOptions())
	bt.Backtrack()
}
