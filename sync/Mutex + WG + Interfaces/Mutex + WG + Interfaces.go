package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mu sync.Mutex
var wg sync.WaitGroup

type store map[string]int

type Customer struct {
	name string
}

type purchase struct {
	customer string
	item     string
	quant    int
}

type Buyer interface {
	Buy(s *store, item string, quant int, purchases chan<- purchase)
}

func (s *store) Buy(item string, quant int) (string, int) {
	q, ok := (*s)[item]
	if !ok {
		fmt.Printf("No such product: %s\n", item)
		return "", 0
	}
	if q < quant {
		fmt.Printf("Not enough %s in stock (requested %d, available %d)\n", item, quant, q)
		return "", 0
	}
	(*s)[item] -= quant
	return item, quant
}

func (c Customer) Buy(s *store, item string, quant int, purchases chan<- purchase) {
	mu.Lock()
	defer mu.Unlock()

	Item, Quant := s.Buy(item, quant)
	if Item != "" && Quant > 0 {
		purchases <- purchase{customer: c.name, item: Item, quant: Quant}
	}
}

func main() {
	str := &store{"apples": 20, "orange": 10, "banana": 5}

	purchases := make(chan purchase, 50)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(2) + 1 // покупок на товар

	customers := []Buyer{
		Customer{"Customer 1"},
		Customer{"Customer 2"},
		Customer{"Customer 3"},
	}

	products := []string{"apples", "orange", "banana"}

	for _, c := range customers {
		for _, p := range products {
			for i := 0; i < n; i++ { // Количество операций = покупателей(3) * товаров(3) * попыток на товар(от 1 до 2)
				wg.Add(1)
				go func(c Buyer, product string) {
					defer wg.Done()
					quant := r.Intn(4) + 1
					c.Buy(str, product, quant, purchases)
				}(c, p)
			}
		}
	}

	wg.Wait()
	close(purchases)

	for p := range purchases {
		fmt.Printf("%s bought %d %s\n", p.customer, p.quant, p.item)
	}

	fmt.Println("\nRemaining stock:", *str)
}
