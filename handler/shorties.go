package handler

import (
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
)

// SaveShorty - метод POST, принимает и сохраняет оригинальный URL
// и возвращает сокращенный
func SaveShorty(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	originalURL := string(bodyBytes)
	shortUrl, err := h.Save(originalURL)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(shortUrl))
}

// GetOriginalURL принимает сокращённый URL и возвращает оригинальный
func GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "shortUrl")
	originalURL, err := h.Get(shortUrl)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(originalURL))
}
