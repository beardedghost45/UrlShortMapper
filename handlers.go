package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

var storage = make(map[string]string)
var topURL = make(map[string]int)
var mu sync.Mutex

type kv struct {
	Key   string
	Value int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetUrl is used for fetching back the original url from short url
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
		if key == url {
			found = true
			originalUrl = value
			break
		}
	}
	if found {
		fmt.Fprintf(w, originalUrl)
	} else {
		fmt.Fprintf(w, "Not Available")
	}

}

// UpdateUrl is used for converting the original url into short url
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

	domain, _ := extractDomain(url)
	if _, exists := topURL[domain]; exists {
		topURL[domain]++
	} else {
		topURL[domain] = 1
	}

	shortCode := generateShortCode()
	mu.Lock()
	storage[shortCode] = url
	mu.Unlock()
	fmt.Fprintf(w, "Shortened URL: http://localhost:8080/url/%s", shortCode)
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

// DomainCount Domain Count method is used to count top 3 domains
func DomainCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var pairs []kv
	mu.Lock()
	for key, value := range topURL {
		pairs = append(pairs, kv{key, value})
	}
	mu.Unlock()

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	limit := 3
	if len(pairs) < 3 {
		limit = len(pairs)
	}

	for i := 0; i < limit; i++ {
		fmt.Fprintf(w, "%s: %d\n", pairs[i].Key, pairs[i].Value)
	}
}

func extractDomain(OriginalURL string) (string, error) {
	schemeEnd := strings.Index(OriginalURL, "://")
	if schemeEnd == -1 {
		return "", errors.New("invalid URL: missing scheme")
	}
	domainStart := schemeEnd + 3
	domainEnd := strings.IndexByte(OriginalURL[domainStart:], '/')
	if domainEnd == -1 {
		return OriginalURL[domainStart:], nil
	}
	return OriginalURL[domainStart : domainStart+domainEnd], nil
}
