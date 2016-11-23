package main

import (
	"fmt"
	"os"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

func main() {
	file, err := os.Open("example.html")

	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(file)

	if err != nil {
		panic(err)
	}

	selector := cascadia.MustCompile("p")

	nodes := selector.MatchAll(doc)

	for _, node := range nodes {
		fmt.Println(node.FirstChild.Data)
	}
}
