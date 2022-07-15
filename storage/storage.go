package storage

import (
	"shorty/storage/heap"
	"shorty/storage/postgres"
)

// StorageInstance - интерфейс хранилища данных, задающий все нужные методы
type StorageInstance interface {
	// Setup настраивает хранилище данных для последующей работы
	Setup()
	// Exists проверяет на наличие записи с идентификатором id в хранилище данных
	Exists(id int64) bool
	// Save сохраняет оригинальный URL в хранилище данных и возвращает сокращенный URL
	Save(url string) (string, error)
	// Get возвращает оригинальный URL из хранилища данных по сокращенному URL
	Get(shortUrl string) (string, error)
}

// GetStorageHandler возвращает нужный объект хранилища данных
// в зависимости от флага storage
func GetStorageHandler(storageMode string) StorageInstance {
	if storageMode == "postgres" {
		return &postgres.PostgresInstance{}
	}
	return &heap.HeapInstance{}
}
