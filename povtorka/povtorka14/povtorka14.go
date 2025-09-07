package main

import (
	"errors"
	"fmt"
)

var (
	ErrItemNotFound   = errors.New("товар не найден")
	ErrOutOfStock     = errors.New("товар закончился")
	ErrNotEnoughStock = errors.New("недостаточно товара на складе")
)

func BuyItem(catalog map[string]int, item string, qty int) error {
	found := false
	for k := range catalog {
		if k == item {
			found = true
		}
	}
	if !found {
		return ErrItemNotFound
	}
	if catalog[item] == 0 {
		return ErrOutOfStock
	}
	if catalog[item] < qty {
		return ErrNotEnoughStock
	}
	return nil
}

func TryBuy(catalog map[string]int, item string, qty int) {
	if err := BuyItem(catalog, item, qty); err != nil {
		switch {
		case errors.Is(err, ErrItemNotFound):
			fmt.Println("товар не найден")
		case errors.Is(err, ErrOutOfStock):
			fmt.Println("товара 0 на складе")
		case errors.Is(err, ErrNotEnoughStock):
			fmt.Println("кол-ва товара не хватает")
		}
	} else {
		fmt.Println("успех!")
	}
}

func main() {
	catalog := map[string]int{"apple": 10, "banana": 0, "orange": 3}

	TryBuy(catalog, "apple", 5)
	TryBuy(catalog, "kakeshki", 5)
	TryBuy(catalog, "orange", 6)
	TryBuy(catalog, "banana", 2)
}
