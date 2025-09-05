package main

import (
	"fmt"
	"sort"
	"strings"
)

type Tel map[string][]string

func AddContact(phoneBook Tel, name string, tel string) {
	tele, _ := phoneBook[name]
	for _, num := range tele {
		if num == tel {
			fmt.Println("tel est")
			return
		}
	}
	phoneBook[name] = append(phoneBook[name], tel)
	fmt.Println("dobavlen")
}

func RemoveContact(phoneBook Tel, name string) {
	if _, ok := phoneBook[name]; ok {
		delete(phoneBook, name)
		fmt.Println("deleted")
		return
	}
	fmt.Println("not found")
}

func Find(phoneBook Tel, name string) {
	if tel, ok := phoneBook[name]; ok {
		fmt.Println("--------------")
		fmt.Printf("Name: %s\n", name)
		for i, t := range tel {
			fmt.Printf("%d) Tel: %s\n", i+1, t)
		}
		fmt.Println("--------------")
		return
	}
	fmt.Println("Not found")
}

func ShowAll(phoneBook Tel) {
	names := make([]string, 0, len(phoneBook))
	for name, _ := range phoneBook {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Println("--------------")
		fmt.Printf("Name: %s\n", name)
		for i, tel := range phoneBook[name] {
			fmt.Printf("%d) Tel: %s\n", i+1, tel)
		}
		fmt.Println("--------------")
	}
}

func Update(phoneBook Tel, name string, ot string, na string) {
	tel, ok := phoneBook[name]
	if !ok {
		fmt.Println("takogo imeni net")
		return
	}
	for i, te := range tel {
		if ot == te {
			tel[i] = na
			phoneBook[name] = tel
			fmt.Println(tel)
			return
		}
	}
	fmt.Println("такого номера у контакта нет")
}

func FindNum(phoneBook Tel, tel string) {
	found := false
	for name, tels := range phoneBook {
		for _, t := range tels {
			if tel == t {
				found = true
				fmt.Println("Контакт", name)
				return
			}
		}
	}
	if !found {
		fmt.Println("кОНТАКТА НЕ НАЙДЕНО")
	}
}

func FindByNamePart(phoneBook Tel, namePart string) {
	found := false
	for name, numbers := range phoneBook {
		if strings.Contains(strings.ToLower(name), strings.ToLower(namePart)) {
			found = true
			fmt.Printf("Name: %s Tel:\n", name)
			for i, t := range numbers {
				fmt.Printf("%d) %s\n", i+1, t)
			}
		}
	}
	if !found {
		fmt.Println("Not found")
	}
}

func AddToFavourite(phoneBook Tel, favourite []string, name string) []string {
	if _, ok := phoneBook[name]; !ok {
		fmt.Printf("Name %s doesn't exist in phone book\n", name)
		return favourite
	}

	for _, nam := range favourite {
		if nam == name {
			fmt.Printf("Name %s is already in favourite\n", name)
			return favourite
		}
	}

	favourite = append(favourite, name)
	fmt.Printf("Name %s added\n", name)
	return favourite
}

func ShowFavourite(favourite []string) {
	sort.Strings(favourite)
	for i, name := range favourite {
		fmt.Printf("%d) Name: %s\n", i+1, name)
	}
}

func RemoveFromFavourite(favourite []string, name string) []string {
	for i, im := range favourite {
		if im == name {
			favourite = append(favourite[:i], favourite[i+1:]...)
			fmt.Printf("Name %s deleted\n", name)
			return favourite
		}
	}
	fmt.Printf("Name %s is not in favourite\n", name)
	return favourite
}

func InsertAt(favourite []string, index int, name string) []string {
	if index < 0 || index > len(favourite) {
		fmt.Println("uncorrect index")
		return favourite
	}
	favourite = append(favourite[:index], (append([]string{name}, favourite[index:]...))...)
	fmt.Println("added")
	return favourite
}

func RemoveAt(favourite []string, index int) []string {
	if index < 0 || index > len(favourite)-1 {
		fmt.Println("uncorrect index")
		return favourite
	}
	favourite = append(favourite[:index], favourite[index+1:]...)
	fmt.Println("removed")
	return favourite
}

func EditAt(favourite []string, index int, name string) []string {
	if index < 0 || index >= len(favourite) {
		fmt.Println("uncorrect index")
		return favourite
	}
	favourite[index] = name
	fmt.Println("edit")
	return favourite
}

func main() {
	book := make(Tel)
	favourite := make([]string, 0, len(book))
	AddContact(book, "Коля", "12345")
	AddContact(book, "Коля", "1234")
	AddContact(book, "кефтеме", "1234")
	AddContact(book, "Аня", "98765")
	favourite = AddToFavourite(book, favourite, "кефтеме")
	favourite = AddToFavourite(book, favourite, "Аня")
	favourite = RemoveFromFavourite(favourite, "Аня")
	favourite = InsertAt(favourite, 0, "чиллер")
	ShowFavourite(favourite)
	Update(book, "Коля", "12345", "чо")
	Find(book, "Коля")
	Find(book, "лариса")
	RemoveContact(book, "Аня")
	ShowAll(book)
	FindNum(book, "чо")
	FindByNamePart(book, "в")
}
