package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/", HandleMain).Methods(http.MethodGet)
	router.HandleFunc("/shorten", HandleShorten).Methods(http.MethodPost)
	router.HandleFunc("/short/{shortKey}", HandleRedirect).Methods(http.MethodGet)

	fmt.Println("URL shortener is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
