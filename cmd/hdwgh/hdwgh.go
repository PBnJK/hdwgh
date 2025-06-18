package main

import (
	"flag"
	"fmt"

	"github.com/pbnjk/hdwgh/pkg/util"
)

func main() {
	flagInput := flag.String("input", "", "Initial article")

	flag.Parse()

	if util.IsValidURL(*flagInput) {
		fmt.Println("TODO")
	}

	fmt.Println("TODO")
}
