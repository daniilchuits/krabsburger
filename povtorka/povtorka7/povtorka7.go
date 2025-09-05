package main

import "fmt"

type Dog struct{}
type Cat struct{}
type Cow struct{}

type Animal interface {
	Sound()
}

func (d Dog) Sound() {
	fmt.Println("гав")
}

func (c Cat) Sound() {
	fmt.Println("мяу")
}

func (co Cow) Sound() {
	fmt.Println("му")
}

func Speak(a Animal) {
	a.Sound()
}

func main() {
	an := []Animal{Dog{}, Cat{}, Cow{}}

	for _, a := range an {
		Speak(a)
	}
}
