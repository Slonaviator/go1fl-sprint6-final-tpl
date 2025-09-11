package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// DetectAndProcess - конвертирует заданную строку в код Морзе и наоборот.
// Возвращает строку и ошибку.
func DetectAndProcess(s string) (string, error) {

	if s == "" {
		return "", errors.New("data is empty")
	}

	if isMorseCode(s) {
		return morse.ToText(s), nil
	}
	return morse.ToMorse(s), nil
}

// isMorseCode - проверяет является ли строка кодом Морзе или нет.
func isMorseCode(s string) bool {
	s = strings.ReplaceAll(s, " ", "") //удаляем пробелы
	for _, char := range s {
		if char != '.' && char != '-' && char != ' ' {
			return false
		}
	}
	return true
}
