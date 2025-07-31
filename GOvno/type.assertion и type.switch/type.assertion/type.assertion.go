package main

import "fmt"

func main() {
	var i interface{} = "hello"
	s, ok := i.(string)
	if ok {
		fmt.Println("Это строка:", s)
	} else {
		fmt.Println("Не строка")
	}
}
