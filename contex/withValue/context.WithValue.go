package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func processRequest(ctx context.Context) {
	defer wg.Done()
	if value := ctx.Value("userID"); value != nil {
		fmt.Println("UserID:", value)
	}
}

func main() {
	ctx := context.WithValue(context.Background(), "userID", 100)

	wg.Add(1)
	go processRequest(ctx)
	wg.Wait()
}
