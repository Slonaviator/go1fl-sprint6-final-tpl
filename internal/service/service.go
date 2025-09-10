package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// DetectAndProcess - конвертирует заданную строку в код Морзе и наоборот.
// Возвращает строку и ошибку.
func DetectAndProcess(s string) (string, error) {

	if isMorseCode(s) {
		return morse.ToText(s), nil
	}
	//if isText(s) {
	return morse.ToMorse(s), nil
	//}
	//return s, fmt.Errorf("conversion error: %w", errors.New("incorrect data"))
}

// isMorseCode - проверяет является ли строка кодом Морзе или нет.
func isMorseCode(s string) bool {
	s = strings.ReplaceAll(s, " ", "") //удаляем пробелы
	for _, char := range s {
		if char != '.' && char != '-' {
			return false
		}
	}
	return true
}

/*// isText - проверяет строку на наличие не буквенных символов.
func isText(s string) bool {
	for _, ch := range s {
		if unicode.IsLetter(ch) {
			return true
		}
	}
	return false
}*/
