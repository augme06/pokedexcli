package repl

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/augme06/pokedexcli/internal/pokeapi"
)

// REPL logic
func StartRepl() error {

	var config pokeapi.Config

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if scanner.Err() != nil {
			log.Fatalf("%v", scanner.Err())
		}

		input := scanner.Text()
		keywords := pokeapi.CleanInput(input)

		var command, parameter string
		switch len(keywords) {
		case 0:
			continue
		case 1:
			command, parameter = keywords[0], ""
		case 2:
			command, parameter = keywords[0], keywords[1]
		}

		if err := Callback(command, parameter, &config); err != nil {
			fmt.Println(err)
		}
	}
}
