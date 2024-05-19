package main

import "github.com/gorilla/mux"

func setupRoutes(r *mux.Router) {
	r.HandleFunc("/url", GetUrl).Methods("GET")
	r.HandleFunc("/url/{string}", UpdateUrl).Methods("POST")
}
