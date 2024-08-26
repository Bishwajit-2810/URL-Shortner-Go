package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	Id           string `json:"id"`
	OriginalURL  string `json:"original_url"`
	ShortlURL    string `json:"short_url"`
	CreationDate string `json:"creation_date"`
}

var urlDB = make(map[string]URL)

func generateShort(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))
	fmt.Println("hasher: ", hasher)
	data := hasher.Sum(nil)
	fmt.Println("hasher data: ", data)
	hash := hex.EncodeToString(data)
	fmt.Println("encoded string: ", hash)
	fmt.Println("final string: ", hash[:8])
	return hash[:8]
}
func createURL(OriginalURL string) string {
	shortURL := generateShort(OriginalURL)
	id := shortURL
	currentDate := time.Now()
	formatted := currentDate.Format("02/01/2006, 3:04:05 PM, Monday")
	urlDB[id] = URL{
		Id:           id,
		OriginalURL:  OriginalURL,
		ShortlURL:    shortURL,
		CreationDate: formatted,
	}
	return shortURL

}
func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")

	}
	return url, nil
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GET method")
}

func ShortlURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	shortUrl_ := createURL(data.URL)
	// fmt.Fprintln(w, shortUrl)
	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: shortUrl_}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	fmt.Println("Starting URL shortner.......")
	defer fmt.Println("Stoping URL shortner........")
	// createURL("https://github.com/Bishwajit-2810")

	// register the handler function to handle all request to the root("/")
	http.HandleFunc("/", handler)
	http.HandleFunc("/shorten", ShortlURLHandler)
	// starting http server on port 3000
	fmt.Println("Starting server on port 3000...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
