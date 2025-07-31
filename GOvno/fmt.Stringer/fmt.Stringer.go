package main

import "fmt"

type Movie struct {
	Title  string
	Year   int
	Rating float64
}

func (m *Movie) String() string {
	return fmt.Sprintf("🎬 Название: %s (Год: %d, Рейтинг: %.1f)", m.Title, m.Year, m.Rating)
}

func main() {
	movies := []Movie{
		{"govni", 1999, 1.23},
		{"elda", 1234, 5.342},
		{"zxc", 228, 10},
	}
	for _, movie := range movies {
		fmt.Println(&movie)
	}
}
