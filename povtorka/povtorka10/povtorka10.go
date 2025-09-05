package main

import "fmt"

type PDF struct{}
type Word struct{}

type Printer interface {
	Print()
}

func (p PDF) Print() {
	fmt.Println("печатаем PDF")
}

func (w Word) Print() {
	fmt.Println("печатаем Word")
}

func main() {
	words := []Printer{PDF{}, Word{}}
	for _, d := range words {
		d.Print()
	}
}
