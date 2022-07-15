package heap

import (
	"errors"
	"math/rand"
	"shorty/base63"
)

type HeapInstance struct {
	IDs         map[int64]bool
	ShortyToUrl map[string]string
}

func (h *HeapInstance) Setup() {
	h.IDs = make(map[int64]bool)
	h.ShortyToUrl = make(map[string]string)
}

func (h *HeapInstance) Exists(id int64) bool {
	return h.IDs[id]
}

func (h *HeapInstance) Save(url string) (string, error) {
	id := rand.Int63()
	for h.Exists(id) {
		id = rand.Int63()
	}
	shortUrl := base63.Encode(id)
	h.ShortyToUrl[shortUrl] = url
	return shortUrl, nil
}

func (h *HeapInstance) Get(shortUrl string) (string, error) {
	v, ok := h.ShortyToUrl[shortUrl]
	if !ok {
		return "", errors.New("shorty not found")
	}
	return v, nil
}
