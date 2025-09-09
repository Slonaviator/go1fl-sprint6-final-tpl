package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Структура сервера
type Server struct {
	Logger *log.Logger
	HTTP   http.Server
}

// NewServer - функция создания сервера
func NewServer(logger *log.Logger) *Server {
	// Создаем роутер
	router := http.NewServeMux()

	// Регистрируем хендлеры
	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	// Создаем структуру сервера
	return &Server{
		Logger: logger,
		HTTP: http.Server{
			Addr:         ":8080",
			Handler:      router,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}

// Start - метод для запуска сервера
func (s *Server) Start() {
	if err := s.HTTP.ListenAndServe(); err != nil {
		s.Logger.Fatal(err)
	}
}
