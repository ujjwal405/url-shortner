package main

import (
	"log"
	"net/http"

	"github.com/ujjwal405/url-shortner/pkg/handlers"
	"github.com/ujjwal405/url-shortner/pkg/store"
)

func main() {

	store := store.NewStore()
	handler := handlers.NewHandlers(store)

	http.Handle("/shorten", handlers.Make(handler.Shorten))
	http.Handle("/shortCode", handlers.Make(handler.ShortCode))

	log.Println("starting the server at 9090")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}
}
