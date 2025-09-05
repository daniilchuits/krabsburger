package main

import "fmt"

func ShowGames(games []string) {
	for _, game := range games {
		fmt.Println(game)
	}
}

func AddGame(games *[]string, newGame string) {
	*games = append(*games, newGame)
}

func Remove(games *[]string, game string) {
	for i, igra := range *games {
		if game == igra {
			*games = append((*games)[:i], (*games)[i+1:]...)
			fmt.Printf("Game %s is deleted\n", game)
			return
		}
	}
	fmt.Println("Not found")
}

func Find(games *[]string, name string) int {
	i := -1
	for i, game := range *games {
		if name == game {
			return i + 1
		}
	}
	return i
}

func Count(games []string) {
	fmt.Println("Length of list", len(games))
}

func main() {
	games := []string{"Minecraft", "Doom", "Witcher", "GTA"}
	ShowGames(games)
	AddGame(&games, "DOTA")
	Remove(&games, "Doom")
	fmt.Println(Find(&games, "GTA"))
	fmt.Println(games)
	Count(games)
}
