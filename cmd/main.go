package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаем логгер
	logger := log.New(
		os.Stdout,     // куда писать
		"[server] ",   // префикс
		log.LstdFlags, // флаги форматирования
	)

	// Создаем сервер
	srv := server.NewServer(logger)
	// Запуск сервера
	srv.Start()
}
