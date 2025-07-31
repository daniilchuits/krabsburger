package main

import (
	"fmt"
	"sync"
)

func main() {
	var rw sync.RWMutex
	var data int
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rw.RLock() // закрывает для других горутин возможность читать, чтобы не было data race
			fmt.Println("read", data)
			rw.RUnlock()
		}()
		wg.Wait()

		wg.Add(1)
		go func() {
			defer wg.Done()
			rw.Lock()
			data = 11
			fmt.Println("read", data)
			rw.Unlock()
		}()
		wg.Wait()
	}
}
