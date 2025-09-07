package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 3)

	go func(ch chan string) {
		for i := 1; i < 6; i++ {
			ch <- fmt.Sprintf("msg%d", i)
		}
		close(ch)
	}(ch)

	for q := range ch {
		time.Sleep(time.Second)
		fmt.Println("Получено:", q)
	}
}
