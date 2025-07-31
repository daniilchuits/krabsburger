package main

import "fmt"

func HandleValue(x any) {
	switch v := x.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", len(v))
	case bool:
		fmt.Println("bool:", v)
	case []int:
		sum := 0
		for _, q := range v {
			sum += q
		}
		fmt.Println("Sum:", sum)
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	HandleValue(42)
	HandleValue("golang")
	HandleValue(true)
	HandleValue([]int{1, 2, 3})
	HandleValue(3.14)
}
