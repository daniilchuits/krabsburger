package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var wg sync.WaitGroup

	init := func() {
		fmt.Println("Инициализация...")
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			once.Do(init)
			fmt.Println("Горутина", id, "запущена")
		}(i)
	}

	wg.Wait()
}
