package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := make(chan float64)
	squares := make(chan float64)
	go func(numbers chan float64) {
		for i := 1; i < 11; i++ {
			numbers <- float64(i)
		}
		close(numbers)
	}(numbers)
	go func(numbers, squares chan float64) {
		for q := range numbers {
			q *= q
			squares <- q
		}
		close(squares)
	}(numbers, squares)
	for z := range squares {
		fmt.Printf("квадрат числа %.0f - %.0f\n", math.Sqrt(z), z)
	}
}
