package main

import (
	"log"
	"net/http"
	"github.com/my-chat-server-api/handler"
)

// main function to boot up everything
func main() {
	peopleHandler := handler.NewPeopleHandler()
	log.Fatal(http.ListenAndServe(":8000", handler.NewRouter(peopleHandler)))
}
