package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var wg sync.WaitGroup

type Config struct {
	AppName string
	Port    int
}

var conf *Config

func LoadConfig() {
	conf = &Config{
		AppName: "AppName",
		Port:    8080,
	}
	fmt.Println("Загрузка конфигурации...")
}

func GetConfig(id int) *Config {
	defer wg.Done()
	once.Do(LoadConfig)
	fmt.Printf("Горутина %d: AppName = %s, Port = %d\n", id, conf.AppName, conf.Port)
	return conf
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go GetConfig(i + 1)
	}
	wg.Wait()
}
