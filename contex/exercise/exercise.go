package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancle()
	temp := make(chan int)
	names := []string{"Dat1", "Dat2", "Dat3"}
	for _, n := range names {
		wg.Add(1)
		go func(name string, ctx context.Context, temp chan<- int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Я датчик %s, я закончился\n", name)
					return
				case <-time.After(500 * time.Millisecond):
					t := rand.Intn(11) + 20
					temp <- t
				}
			}
		}(n, ctx, temp)
	}
	go func(temp chan int) {
		for v := range temp {
			fmt.Println("Temp:", v)
		}
	}(temp)
	wg.Wait()
	// я не знаю почему, но я делал это прям ужасно долго, меня выбесило все, я полностью переделывал несколько
	// десятков раз. Я В АХУЕ
	close(temp)
}
