package storage

import (
	"shorty/storage/heap"
	"shorty/storage/postgres"
)

type StorageHandler interface {
	Exists(id int64) bool
	Setup()
	Save(url string) (string, error)
	Get(shortUrl string) (string, error)
}

func GetStorageHandler(mode string) StorageHandler {
	if mode == "postgres" {
		return &postgres.PostgresHandler{}
	}
	return &heap.HeapHandler{}
}
