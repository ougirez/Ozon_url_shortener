package main

import (
	"flag"
	"math/rand"
	"net/http"
	"shorty/handler"
	"time"
)

func main() {
	// чтобы при перезапуске сервера пакет rand генерировал новые значения для ID
	rand.Seed(time.Now().UnixNano())
	// парсим флаг -storage, который определяет, какое хранилище данных использует
	// приложение: postgres или heap (память приложения)
	storageMode := flag.String("storage", "", "postgres or heap")
	flag.Parse()
	r := handler.NewHandler(*storageMode)
	http.ListenAndServe(":3000", r)
}
