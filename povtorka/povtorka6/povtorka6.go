package main

import "fmt"

type Car struct{}
type Bike struct{}
type Plane struct{}

func (c Car) Move() {
	fmt.Println("Машина едет по дороге")
}

func (b Bike) Move() {
	fmt.Println("велик едет по тратуару")
}

func (p Plane) Move() {
	fmt.Println("самолет летит по небу")
}

type Transport interface {
	Move()
}

func Start(t Transport) {
	t.Move()
}

func main() {
	tran := []Transport{Car{}, Bike{}, Plane{}}
	for _, t := range tran {
		Start(t)
	}
}
