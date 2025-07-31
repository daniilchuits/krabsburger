package main

import (
	"fmt"
	"reflect"
)

// type User struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	u := User{Name: "Alice", Age: 30}
// 	t := reflect.TypeOf(u)
// 	v := reflect.ValueOf(u)

// 	for i := 0; i < t.NumField(); i++ {
// 		field := t.Field(i)
// 		value := v.Field(i)
// 		fmt.Printf("Поле: %s, Тип: %s, Значение: %v\n", field.Name, field.Type, value)
// 	}
// }

// func main() {
// 	x := 10
// 	v := reflect.ValueOf(&x).Elem()
// 	fmt.Println("До:", x)

// 	if v.CanSet() { // SetInt - метод из пакета reflect, который меняет значение int, int8 и тд
// 		v.SetInt(99) // а методом CanSet мы проверяем можно ли изменить переменную v
// 	}

// 	fmt.Println("После:", x)
// }

// func Inspect(i interface{}) {
// 	v := reflect.ValueOf(i)
// 	t := reflect.TypeOf(i)

// 	fmt.Println("Тип:", t)
// 	fmt.Println("Значение:", v)
// 	fmt.Println("Kind:", v.Kind())

// 	switch v.Kind() {
// 	case reflect.Int:
// 		fmt.Println("Целое число:", v.Int())
// 	case reflect.String:
// 		fmt.Println("Строка:", v.String())
// 	}
// }

// func main() {
// 	Inspect([2]int{1, 2})
// }

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func zxc() {
	t := reflect.TypeOf(Product{})
	fmt.Println(t)
}

func main() {
	zxc()
}
