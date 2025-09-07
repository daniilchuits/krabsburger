package main

import (
	"errors"
	"fmt"
	"strings"
)

var ErrBookNotFound = errors.New("книга не найдена")

func FindBook(library []string, title string) (string, error) {
	for _, book := range library {
		if strings.EqualFold(book, title) {
			return title, nil
		}
	}
	return "", ErrBookNotFound
}

func main() {
	Lib := []string{"Chingiz", "Gomunkul"}
	a, err := FindBook(Lib, "gomunkul")
	if err != nil {
		if errors.Is(err, ErrBookNotFound) {
			fmt.Println("kniga ne naydena")
		} else {
			fmt.Println("error:", err)
		}
		return
	}
	fmt.Printf("kniga %s found\n", a)
}
