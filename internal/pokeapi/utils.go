package pokeapi

import (
	"fmt"
	"math/rand"
	"strings"
)

func CleanInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}

func CatchTry(exp int) bool {
	chance := rand.Intn(exp)
	return chance < exp/2
}

func AddToPokedex(pokemon Pokemon) {
	Pokedex[pokemon.Name] = pokemon
}

func FormatOutput(pokemon Pokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", s.StatInfo.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Type {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
}
