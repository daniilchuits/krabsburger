package main

import (
	"errors"
	"fmt"
)

type InvalidAgeError struct {
	Age int
}

func (v *InvalidAgeError) Error() string {
	return "Age error: it must be 18 or more"
}

func registerUser(age int) error {
	if age < 18 {
		fmt.Println(&InvalidAgeError{Age: age})
		return &InvalidAgeError{Age: age}
	}
	fmt.Println("User registred")
	return nil
}

func main() {
	err := registerUser(-5)

	var iAE *InvalidAgeError
	if errors.As(err, &iAE) {
		fmt.Println("Cannot register")
	} else if err != nil {
		fmt.Println("Неизвестная ошибка:", err)
	} else {
		fmt.Println("Регистрация прошла успешно")
	}
}
