package main

import (
	"errors"
	"fmt"
)

var ErrNoDBName = errors.New("имя БД не задано")
var ErrNoConnection = errors.New("нет подключения к серверу")

func connectionToDB(name string) (string, error) {
	if name == "" {
		return "", ErrNoDBName
	} else if name == "offline" {
		return "", ErrNoConnection
	}
	return fmt.Sprintf("Connection with DB: %s \n", name), nil
}

func startApp(name string) (string, error) {
	result, err := connectionToDB(name)
	if err != nil {
		return "", fmt.Errorf("ошибка запуска: %w", err)
	}
	return fmt.Sprint("Succssesful connection\n", result), nil
}

func main() {
	for {
		var name string
		fmt.Println("Write the name of your DB")
		fmt.Scanln(&name)

		db, err := startApp(name)
		if errors.Is(err, ErrNoDBName) {
			fmt.Println("Error: undefind DB name")
		} else if errors.Is(err, ErrNoConnection) {
			fmt.Println("Error: no connection with server")
		} else {
			fmt.Print("App is sucsessfuly launched\n", db)
		}
	}
}
