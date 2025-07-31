package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

type Book struct {
	Title  string
	Author string
	Genre  string
}

type Library struct {
	Books []Book
}

func (l *Library) ShowAll() {
	for i, book := range l.Books {
		fmt.Printf("%d. %s by %s, genre [%s]\n", i+1, book.Title, book.Author, book.Genre)
	}
}

func (l Library) ShowByGenre(gen string) {
	gen = strings.ToLower(gen)
	found := false
	for _, book := range l.Books {
		if strings.ToLower(book.Genre) == gen {
			fmt.Printf("Genre: [%s]\nTitle: %s\nAuthor: %s\n \n", book.Genre, book.Title, book.Author)
			found = true
		}
	}
	if !found {
		fmt.Println("No books of this genre\n")
	}
}

func (l *Library) AddBook(Book Book) {
	for _, boo := range l.Books {
		if boo.Title == Book.Title {
			fmt.Println("Book is already available")
			return
		}
	}
	l.Books = append(l.Books, Book)
}

func (l *Library) DeleteByTitle(name string) {
	name = strings.ToLower(name)
	for i, book := range l.Books {
		if strings.ToLower(book.Title) == name {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			fmt.Println("Book is deleted")
			return
		}
	}
	fmt.Println("Book is not found")
}

func (l Library) SearchByAuthor(auth string) {
	found := false
	for _, book := range l.Books {
		if strings.ToLower(book.Author) == auth {
			found = true
		}
	}
	if !found {
		for _, book := range l.Books {
			fmt.Printf("Books of %s are not found\n", book.Author)
			break
		}
		fmt.Println()
	}
	if found {
		fmt.Println("Books of author ", auth)
		for _, book := range l.Books {
			if strings.ToLower(book.Author) == auth {
				fmt.Printf("%s\n", book.Title)
			}
		}
		fmt.Println()
	}
}

func (l Library) RecommendRandomBookByGenre(genre string) {
	genre = strings.ToLower(genre)
	var filteredBooks []Book

	for _, book := range l.Books {
		if strings.ToLower(book.Genre) == genre {
			filteredBooks = append(filteredBooks, book)
		}
	}

	if len(filteredBooks) == 0 {
		fmt.Println("No books of this genre found")
		return
	}

	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	r := randGen.Intn(len(filteredBooks))
	BOOK := filteredBooks[r]
	fmt.Printf("Book of this genre: %s by %s\n", BOOK.Title, BOOK.Author)
}

func (l Library) KeyWords() {
	fmt.Println("Enter key words (Enter b to stop)")
	reader := bufio.NewReader(os.Stdin)
	var msges []string
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
		}
		if strings.TrimSpace(strings.ToLower(msg)) == "b" {
			break
		}
		msg = strings.TrimSpace(msg)
		msges = append(msges, msg)
	}

	var wg sync.WaitGroup
	var found bool

	for _, key := range msges {
		for _, book := range l.Books {
			wg.Add(1)
			go func(key string, book Book) {
				defer wg.Done()
				if strings.EqualFold(key, book.Author) || strings.EqualFold(key, book.Genre) || strings.EqualFold(key, book.Title) {
					found = true
					fmt.Printf("Book is found\nTitle: %s\nAuthor: %s\nGenre: %s\n", book.Title, book.Author, book.Genre)
				}

			}(key, book)
		}
	}
	wg.Wait()
	if !found {
		fmt.Println("Book is not found")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	lib := Library{
		Books: []Book{
			{"Kazaki razboyniki", "Yuriy Gargarin", "Fantasy"},
			{"Velikoe poedanie kozyavok", "Temniy prince", "Poem"},
			{"Polniy bak", "Buster", "Science fiction"},
			{"Kormleniye detey", "Papich", "Real story"},
			{"Samorazvitiye", "Arsen Markaryan", "Fantasy"},
		},
	}
	fmt.Println("Commands\nShow all - s\nShow all for your genre - g\n" +
		"Add book - a\nDelete a book - d\nSearch by author - SBA\nEnd program - b\n" +
		"Recomend by genre - r\nFind by key words - k")
	for {
		fmt.Println("Enter the desired action")
		answer, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		answer = strings.TrimSpace(answer)
		answer = strings.ToLower(answer)
		fmt.Println()

		switch answer {
		case "s":
			lib.ShowAll()

		case "g":
			fmt.Println("Enter a genre")
			gen, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			gen = strings.TrimSpace(gen)
			lib.ShowByGenre(gen)

		case "b":
			fmt.Println("Bye bye")
			return // или break, если цикл

		case "a":
			fmt.Println("Enter a title for book")
			titl, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			titl = strings.TrimSpace(titl)

			fmt.Println("Enter an author")
			aut, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			aut = strings.TrimSpace(aut)

			fmt.Println("Enter a genre")
			ge, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			ge = strings.TrimSpace(ge)

			fmt.Println("Book is added")
			lib.AddBook(Book{
				Title:  titl,
				Author: aut,
				Genre:  ge,
			})

		case "d":
			fmt.Println("Enter a book title")
			BT, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			BT = strings.TrimSpace(strings.ToLower(BT))
			lib.DeleteByTitle(BT)

		case "sba":
			fmt.Println("Enter an author for search")
			aut, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			aut = strings.TrimSpace(strings.ToLower(aut))
			lib.SearchByAuthor(aut)

		case "r":
			fmt.Println("Enter a genre")
			genr, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			genr = strings.TrimSpace(strings.ToLower(genr))
			lib.RecommendRandomBookByGenre(genr)

		case "k":
			lib.KeyWords()

		default:
			fmt.Println("No such an action")
		}
	}

}
