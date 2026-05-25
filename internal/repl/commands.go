package repl

import (
	"fmt"
	"os"

	"github.com/augme06/pokedexcli/internal/ansi"
	"github.com/augme06/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(parameter string, config *pokeapi.Config) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists all pokemon in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects a pokemon, if caught",
			callback:    commandInspect,
		},
		"clear": {
			name:        "clear",
			description: "Clears all content on screen",
			callback:    commandClear,
		},
	}
}

func Callback(command, parameter string, config *pokeapi.Config) error {
	if _, ok := GetCommands()[command]; !ok {
		return fmt.Errorf("Unknown command")
	}
	if err := GetCommands()[command].callback(parameter, config); err != nil {
		return err
	}
	return nil
}

// Callbacks
func commandExit(parameter string, config *pokeapi.Config) error {
	fmt.Print(ansi.ExitAltScreen)
	fmt.Printf("You have captured %d Pokemon(s)!\n", len(pokeapi.Pokedex))
	os.Exit(0)
	return nil
}

func commandHelp(parameter string, config *pokeapi.Config) error {
	commands := GetCommands()
	fmt.Println(ansi.Format("Usage:", ansi.Bold))
	for _, c := range commands {
		fmt.Printf("%s: %s\n", ansi.Format(c.name, ansi.Green), c.description)
	}
	return nil
}

func commandMap(parameter string, config *pokeapi.Config) error {
	return pokeapi.GetNextMap(config)
}

func commandMapb(parameter string, config *pokeapi.Config) error {
	return pokeapi.GetPreviousMap(config)
}

func commandExplore(parameter string, config *pokeapi.Config) error {
	return pokeapi.Explore(parameter)
}

func commandCatch(parameter string, config *pokeapi.Config) error {

	pokemon, err := pokeapi.GetPokemon(parameter)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", ansi.Format(pokemon.Name, ansi.Bold, ansi.Yellow))
	if pokeapi.CatchTry(pokemon.BaseExperience) {
		fmt.Printf(ansi.Format("%s was caught!\n", ansi.Bold), pokemon.Name)
		pokeapi.AddToPokedex(pokemon)
	} else {
		fmt.Printf(ansi.Format("%s escaped!\n", ansi.Bold), pokemon.Name)
	}
	return nil
}

func commandInspect(parameter string, config *pokeapi.Config) error {
	info, ok := pokeapi.Pokedex[parameter]
	if !ok {
		fmt.Println("pokemon not caught")
		return nil
	}

	pokeapi.FormatOutput(info)

	return nil
}

func commandClear(parameter string, config *pokeapi.Config) error {
	fmt.Print(ansi.Clear)
	return nil
}
