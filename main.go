package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
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

func main() {
	fmt.Println("Starting URL shortner.......")
	defer fmt.Println("Stoping URL shortner........")
	createURL("https://github.com/Bishwajit-2810")

}
