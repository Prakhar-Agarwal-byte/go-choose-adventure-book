package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prakhar-agarwal-byte/go-choose-adventure-book/models"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the server on")
	fileName := flag.String("json", "gopher.json", "json file containing the story")
	flag.Parse()
	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	s, err := models.JsonStory(f)
	if err != nil {
		panic(err)
	}
	h := models.NewStoryHandler(s)
	fmt.Printf("Starting the server on port: %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}