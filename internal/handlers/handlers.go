package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// HomeHandler обрабатывает корневой эндпоинт
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	// Читаем HTML-файл
	file, err := os.ReadFile("../index.html")
	if err != nil {
		log.Printf("reading error index.html: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок и отправляем ответ
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(file)
}

// UploadHandler обрабатывает загрузку файла
func UploadHandler(w http.ResponseWriter, r *http.Request) {

	// Парсим форму
	err := r.ParseMultipartForm(32 << 20) // 32 MB limit
	if err != nil {
		log.Printf("form parsing error: %v", err)
		http.Error(w, "file upload error", http.StatusInternalServerError)
		return
	}

	// Получаем файл из формы
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Printf("file receipt error: %v", err)
		http.Error(w, "the file was not found in the form", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Чтение файла
	content, err := io.ReadAll(file)
	if err != nil {
		log.Printf("file reading error: %v", err)
		http.Error(w, "file reading error", http.StatusInternalServerError)
		return
	}

	// Конвертация из пакета service
	convertedText, err := service.DetectAndProcess(string(content))
	if err != nil {
		log.Printf("conversion error: %v", err)
		http.Error(w, "conversion error", http.StatusInternalServerError)
		return
	}

	// Формируем имя файла
	timeFileName := time.Now().UTC().Format("2006-01-02_15-04-05") //time.Now().UTC().String()
	ext := filepath.Ext(handler.Filename)
	newFileName := fmt.Sprintf("convertedFile_%s%s", timeFileName, ext)

	// Записываем результат в файл
	fileNew, err := os.Create(newFileName)
	if err != nil {
		log.Printf("error writing to a file: %v", err)
		http.Error(w, "error recording the result", http.StatusInternalServerError)
		return
	}
	defer fileNew.Close()

	_, err = fileNew.WriteString(convertedText)
	if err != nil {
		log.Printf("error writing to a file: %v", err)
		http.Error(w, "error recording the result", http.StatusInternalServerError)
		return
	}
	// Возвращаем результат
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"result":   convertedText,
		"filename": newFileName,
	})
}
