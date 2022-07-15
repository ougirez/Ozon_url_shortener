package handler

import (
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
)

// SaveShorty :  принимает и сохраняет оригинальный URL и возвращает сокращенный
// URL : /
// Method : POST
// Example Body : https://github.com/ougirez/ozon_URL_shortener
// Example Output : http://localhost:3000/CG5Xmv7TpM
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
// URL : /{shortURL}
// Method : GET
// Output : текст с оригинальным URL. Решил не использовать JSON, так как
// ответ совсем простой, к тому же в дальнейшем скорее всего исползуется для редиректа,
// а не для какого-либо представления на фронте.
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
