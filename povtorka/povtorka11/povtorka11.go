package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var ErrFileNotExist = errors.New("файл не существует")
var ErrInvalidExt = errors.New("неверное расширение файла")

func ReadConfig(path string) (string, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", ErrFileNotExist
		}
		return "", fmt.Errorf("ошибка доступа к файлу %w", err)
	}

	if filepath.Ext(info.Name()) != ".conf" {
		return "", ErrInvalidExt
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения файла %w", err)
	}

	return string(data), nil
}

func main() {
	files := []string{"config.conf", "config.txt", "nofile.conf"}

	for _, f := range files {
		fmt.Println("=== Проверка файла:", f, "===")
		content, err := ReadConfig(f)
		if err != nil {
			if errors.Is(err, ErrFileNotExist) {
				fmt.Println("Ошибка: файл не найден")
			} else if errors.Is(err, ErrInvalidExt) {
				fmt.Println("Ошибка: неверное расширение")
			} else {
				fmt.Println("Error:", err)
				fmt.Println("Внутренняя ошибка:", errors.Unwrap(err))
			}
			continue
		}
		fmt.Println("Содержимое файла:", content)
	}
}
