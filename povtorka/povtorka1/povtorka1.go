package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func (u User) SayHello() {
	fmt.Println("Hi", u.Name)
}

func (u User) IsAdult() bool {
	if u.Age >= 18 {
		return true
	}
	return false
}

func IncreaseAge(u *User, years int) {
	u.Age += years
	fmt.Println("Age is", u.Age)
}

func main() {
	Users := map[string]User{
		"Anna": {"Anna", 16},
		"Bob":  {"Bob", 100},
		"Alex": {"Alex", 22},
	}

	for _, us := range Users {
		us.SayHello()
		IsAdult := us.IsAdult()
		if IsAdult {
			fmt.Println("I'm adult")
		} else {
			fmt.Println("I'm not adult")
		}
	}

	for name, u := range Users {
		IncreaseAge(&u, 2)
		Users[name] = u // обязательно перезаписывать в карту измененные значения
		fmt.Println(u.Age)
	}

	delete(Users, "Anna")
	fmt.Println(Users)
	fmt.Println()

	if u, ok := Users["Alex"]; ok {
		fmt.Println("Found", u)
	} else {
		fmt.Println("Not found(")
	}
	fmt.Println()

	sum := 0
	for _, u := range Users {
		sum += u.Age
	}
	fmt.Println("Sum =", sum)
}
