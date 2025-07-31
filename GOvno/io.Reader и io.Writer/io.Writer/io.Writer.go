package main

import (
	"fmt"
	"os"
)

func main() {
	writer := os.Stdout // реализует io.Writer
	data := []byte("Привет, Writer!\n")
	n, err := writer.Write(data)
	if err != nil {
		fmt.Println("Ошибка записи:", err)
		return
	}
	fmt.Printf("Записано %d байт\n", n)
}
