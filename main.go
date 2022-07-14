package main

import (
	"flag"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"io/ioutil"
	"math/rand"
	"net/http"
	"shorty/storage"
	"time"
)

var h storage.StorageHandler

func SaveShorty(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}
	originalURL := string(bodyBytes)
	shortUrl, err := h.Save(originalURL)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(200)
	w.Write([]byte(shortUrl))
}

func GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "shortUrl")
	originalURL, err := h.Get(shortUrl)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(200)
	w.Write([]byte(originalURL))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	storageMode := flag.String("storage", "", "postgres or heap")
	flag.Parse()
	h = storage.GetStorageHandler(*storageMode)
	h.Setup()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/", SaveShorty)
	r.Get("/{shortUrl}", GetOriginalURL)
	http.ListenAndServe(":3000", r)
}
