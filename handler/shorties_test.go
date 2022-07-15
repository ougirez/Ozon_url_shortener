package handler

import (
	"bytes"
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"shorty/base63"
	"shorty/storage/heap"
	"strconv"
	"strings"
	"testing"
)

func TestSaveShorty(t *testing.T) {
	h = &heap.HeapInstance{}
	h.Setup()
	var str = []byte("https://www.youtube.com/watch?v=a3Adgew7q5g")

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(str))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := chi.NewRouter()
	handler.Post("/", SaveShorty)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	responseBody := strings.TrimSpace(rr.Body.String())
	shortUrl := strings.TrimPrefix(responseBody, "http://localhost:3000/")
	if len(shortUrl) != 10 {
		t.Errorf("incorrect length of shor URL: got %d want 10", len(shortUrl))
	}
	for c := range shortUrl {
		if !strings.Contains(base63.ALPHABET, strconv.Itoa(c)) {
			t.Errorf("unexpected character in handler returned body: %c", c)
		}
	}
}

func TestGetOriginalURL(t *testing.T) {
	h = &heap.HeapInstance{
		IDs:         map[int64]bool{512351235: true},
		ShortyToUrl: map[string]string{"27I5tQ6N4E": "https://github.com/ougirez/ozon_URL_shortener"},
	}
	req, err := http.NewRequest("GET", "/27I5tQ6N4E", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := chi.NewRouter()
	handler.Get("/{shortUrl}", GetOriginalURL)
	handler.ServeHTTP(rr, req)
	expected := "https://github.com/ougirez/ozon_URL_shortener"
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetOriginalURLNotFound(t *testing.T) {
	h = &heap.HeapInstance{
		IDs:         map[int64]bool{512351235: true},
		ShortyToUrl: map[string]string{"27I5tQ6N4E": "https://github.com/ougirez/ozon_URL_shortener"},
	}
	req, err := http.NewRequest("GET", "/SbvZLENe0", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := chi.NewRouter()
	handler.Get("/{shortUrl}", GetOriginalURL)
	handler.ServeHTTP(rr, req)
	expected := "shorty not found"
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
