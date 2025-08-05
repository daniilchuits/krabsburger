package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Время загрузок для каждого "сервера"
	delays := []time.Duration{
		1 * time.Second, // успеет до таймаута
		3 * time.Second, // не успеет
		5 * time.Second, // не успеет
	}

	var wg sync.WaitGroup

	for i, d := range delays {
		wg.Add(1)

		// Контекст с таймаутом 2 сек
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

		go func(ctx context.Context, d time.Duration, i int) {
			defer wg.Done()
			defer cancel()

			fmt.Printf("Server %d: downloading started\n", i+1)

			// Канал для сигнала об успешном завершении
			done := make(chan struct{})

			// Запускаем "загрузку"
			go func() {
				time.Sleep(d) // эмуляция времени скачивания
				close(done)   // сигнал, что загрузка завершена
			}()

			// Ждём либо завершения, либо таймаута
			select {
			case <-done:
				fmt.Printf("Server %d: downloading ended successfully ✅\n", i+1)
			case <-ctx.Done():
				if ctx.Err() == context.DeadlineExceeded {
					fmt.Printf("Server %d: downloading canceled by timeout ⏱\n", i+1)
				} else if ctx.Err() == context.Canceled {
					fmt.Printf("Server %d: downloading canceled manually ❌\n", i+1)
				}
			}
		}(ctx, d, i)
	}

	wg.Wait()
}
