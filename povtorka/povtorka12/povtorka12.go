package main

import (
	"errors"
	"fmt"
	"strings"
)

var ErrEmptyUsername = errors.New("пустое имя пользователя")
var ErrInvalidEmail = errors.New("недопустимый имейл")
var ErrWeakPassword = errors.New("короткий пароль")

func ValidateUsername(username, email, password string) error {
	var a, b, c error
	if len(username) == 0 {
		a = ErrEmptyUsername
	}
	if !strings.Contains(email, "@") {
		b = fmt.Errorf("Проверка имейла: %w", ErrInvalidEmail)
	}
	if len(password) < 6 {
		c = ErrWeakPassword
	}
	return errors.Join(a, b, c)
}

func main() {
	us := [][]string{
		{"ea", "zxc@zxc", "123456789"},
		{"qwe", "qwe", "qweasdzxc"},
		{"", "asd", "qwe"},
	}
	for _, u := range us {
		fmt.Println("Проверка")
		err := ValidateUsername(u[0], u[1], u[2])
		if err != nil {
			if errors.Is(err, ErrEmptyUsername) {
				fmt.Println("-", ErrEmptyUsername)
			}
			if errors.Is(err, ErrInvalidEmail) {
				fmt.Println("-", ErrInvalidEmail)
			}
			if errors.Is(err, ErrWeakPassword) {
				fmt.Println("-", ErrWeakPassword)
			}
			fmt.Println("Все ошибки:", err)
		} else {
			fmt.Println("Все норм")
		}
		fmt.Println()
	}
}
