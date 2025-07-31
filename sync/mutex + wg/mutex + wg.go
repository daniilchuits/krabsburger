package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1) // каждый раз (1000 раз) добавляем 1 к ожиданию для завершения всех горутин
		go func() {
			defer wg.Done() // отнимаем от этой 1000 единицу каждый раз, когда горутина выполняется
			mu.Lock()       // блокируем доступ, чтобы 1 горутина могла читать-писать данные
			counter++
			mu.Unlock() // наша горутина сделала свое дело, открываем доступ
		}()
	}
	wg.Wait() // ожидаем выполнения всех 1000 горутин и после этого переходим к след действию
	fmt.Println(counter)
}
