package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"shorty/storage"
)

// переменная, обеспечивающая доступ к хранилищу данных
var h storage.StorageInstance

func NewHandler(storageMode string) chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/", SaveShorty)
	router.Get("/{shortUrl}", GetOriginalURL)
	h = storage.GetStorageHandler(storageMode)
	h.Setup()
	return router
}
