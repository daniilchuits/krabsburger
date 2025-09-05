package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func rabota(mu *sync.Mutex, total time.Duration, ctx, ctxq context.Context) {
	timeout:=time.After(500 * time.Millisecond)
	total:=0
	go func() {
		for i:=0; i < 5; i++{
			fmt.Printf("[%d] Начало загрузки file%d...\n", i + 1, i + 1)
			Zagr:=time.Duration(rand.Intn(5)+1) * time.Second
			mu.Lock()
			total+=time.Duration(Zagr)
			mu.Unlock()
			ctxq, cancle:=context.WithCancel(context.Background())
			for{
				select{
				case total > 7 * time.Second:
					cancle()
				case <- ctxq.Done():
					fmt.Printf("Загрузка завершена[%d]\n", i + 1)
				}
			}
		}
	}
}

func main() {
	ctx, cancle:=context.WithTimeout(context.Background(), 7 * time.Second)
	defer cancle()
	var totalTime time.Duration
	var mu sync.Mutex

	go rabota(mu, totalTime,ctx, ctxq)
}
