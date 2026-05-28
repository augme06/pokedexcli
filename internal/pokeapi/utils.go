package pokeapi

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/augme06/pokedexcli/internal/ansi"
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
	fmt.Println("Name: " + ansi.Format(pokemon.Name, ansi.Yellow))
	fmt.Println("Height: " + fmt.Sprintf(ansi.Format("%d", ansi.Yellow), pokemon.Height))
	fmt.Println("Weight: " + fmt.Sprintf(ansi.Format("%d", ansi.Yellow), pokemon.Weight))
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  -%s: ", s.StatInfo.Name)
		fmt.Printf(ansi.Format("%d\n", ansi.Blue, ansi.Bold), s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Type {
		fmt.Printf(ansi.Format("  - %s\n", ansi.Bold), t.Type.Name)
	}
}
