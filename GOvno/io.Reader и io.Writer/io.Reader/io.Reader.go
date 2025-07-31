package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello, Reader!")

	buf := make([]byte, 5) // читаем по 5 байт за раз
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Printf("Прочитано %d байт: %q\n", n, buf[:n])
	}
}
