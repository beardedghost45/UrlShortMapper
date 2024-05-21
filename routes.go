package main

import "github.com/gorilla/mux"

func setupRoutes(r *mux.Router) {
	r.HandleFunc("/url/{ShortURL}", GetUrl).Methods("GET")
	r.HandleFunc("/url", UpdateUrl).Methods("POST")
	r.HandleFunc("/domain", DomainCount).Methods("GET")
}
