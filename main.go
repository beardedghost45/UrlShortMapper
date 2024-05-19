package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	setupRoutes(r)
	log.Fatal(http.ListenAndServe(":8000", r))

}
