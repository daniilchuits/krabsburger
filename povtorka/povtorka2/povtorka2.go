package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

type Library struct {
	Books map[string]Book
}

func (l *Library) AddBook(b Book) {
	for name, _ := range l.Books {
		if name == b.Title {
			fmt.Println("Suschestvuet")
			return
		}
	}
	l.Books[b.Title] = b
	fmt.Println("Added", b)
}

func (l *Library) Remove(title string) {
	found := false
	for name, _ := range l.Books {
		if title == name {
			delete(l.Books, title)
			found = true
			fmt.Println("Removed")
			return
		}
	}
	if !found {
		fmt.Println("Not found")
	}
}

func (l Library) Find(title string) (Book, bool) {
	book, ok := l.Books[title]
	return book, ok
}

func (l Library) All() {
	for _, book := range l.Books {
		fmt.Println(book)
	}
}

func (l Library) TotalPages() {
	total := 0
	for _, book := range l.Books {
		total += book.Pages
	}
	fmt.Println("Total pages:", total)
}

func main() {
	Lib := Library{
		Books: map[string]Book{
			"Govnou": {"Govnou", "Kolya", 100},
			"Kishki": {"Kishki", "Sasavot", 234},
		},
	}

	Lib.AddBook(Book{"Piter", "Grifin", 444})
	Lib.Remove("Govnou")
	Lib.All()
	fmt.Println()
	if book, ok := Lib.Find("Piter"); ok {
		fmt.Println(book)
	} else {
		fmt.Println("Not")
	}
	if bok, ok := Lib.Find("Kormleniye"); ok {
		fmt.Println(bok)
	} else {
		fmt.Println("Not")
	}
	Lib.TotalPages()
}
