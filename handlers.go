package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var storage = make(map[string]string)
var mu sync.Mutex

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)
	var found = false
	var originalUrl string
	url := vars["ShortURL"]

	for key, value := range storage {
		if value == url {
			found = true
			originalUrl = key
			break
		}
	}
	if found {
		fmt.Print(originalUrl)
	} else {
		fmt.Print("Not Available")
	}

}

func UpdateUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortCode := generateShortCode()
	mu.Lock()
	storage[shortCode] = url
	mu.Unlock()
	fmt.Fprintf(w, "Shortened URL: http://localhost:8080/%s", shortCode)
}

func generateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for {
		shortCode := make([]byte, 6)
		for i := range shortCode {
			shortCode[i] = charset[rand.Intn(len(charset))]
		}
		code := string(shortCode)
		mu.Lock()
		if _, exists := storage[code]; !exists {
			mu.Unlock()
			return code
		}
		mu.Unlock()
	}
}
