package main

import (
	"github.com/gorilla/handlers"

	"log"
	"net/http"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("serving %s", dir)
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, http.FileServer(http.Dir(dir))))
	http.ListenAndServe(":8888", nil)
}
