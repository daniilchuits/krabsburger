package main

import "fmt"

func doSomething(x any) {
	fmt.Println("Значение:", x)
}

func printAnything(val interface{}) {
	fmt.Println("Получено значение:", val)
}

type Book struct {
	Name   string
	Author string
}

func main() {
	b := Book{Name: "Polevih", Author: "Cvetov"}
	printAnything(43)
	printAnything(b)
	doSomething("hello")
	doSomething([]int{1, 2, 3})
}
