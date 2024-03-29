package main

import (
	cy "cyoa/story"
	"log"
	"net/http"

	//"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the cyoa web application on")
	filename := flag.String("file", "../../gopher.json", "the JSON file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cy.JsonStory(f)
	if err != nil {
		panic(err)
	}

	h := cy.NewHandler(story)

	fmt.Printf("Starting the server on %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
