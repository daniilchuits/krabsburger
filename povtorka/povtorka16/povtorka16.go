package main

import "fmt"

func main() {
	ch := make(chan int)
	go func(ch chan int) {
		for i := 1; i < 6; i++ {
			ch <- i
		}
		close(ch)
	}(ch)
	for q := range ch {
		fmt.Println("Получено:", q)
	}
}
