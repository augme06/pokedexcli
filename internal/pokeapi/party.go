package pokeapi

import (
	"fmt"

	"github.com/augme06/pokedexcli/internal/ansi"
)

var party = make(map[string]Pokemon, 6)

func GetParty() {
	if len(party) == 0 {
		fmt.Println("Your party is empty!")
		return
	}

	fmt.Println("YOUR PARTY")
	for _, p := range party {
		fmt.Printf(ansi.Format(" - %s\n", ansi.Yellow), p.Name)
	}
}

func AddToParty(name string) error {
	if len(party) == 6 {
		return fmt.Errorf("Party is full!")
	}

	p, ok := Pokedex[name]
	if !ok {
		return fmt.Errorf("You don't have have that Pokemon!")
	}

	if _, ok := party[name]; ok {
		return fmt.Errorf("Pokemon already in the party!")
	}

	party[name] = p
	return nil
}

func RemoveFromParty(name string) error {
	_, ok := party[name]
	if !ok {
		return fmt.Errorf("Pokemon not found!")
	}

	delete(party, name)
	return nil
}
