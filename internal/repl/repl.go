package repl

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/augme06/pokedexcli/internal/ansi"
	"github.com/augme06/pokedexcli/internal/pokeapi"
)

func printBanner() {
	fmt.Println(ansi.Format(`
▄██   ▄      ▄████████    ▄███████▄ 
███   ██▄   ███    ███   ███    ███ 
███▄▄▄███   ███    █▀    ███    ███ 
▀▀▀▀▀▀███  ▄███▄▄▄       ███    ███ 
▄██   ███ ▀▀███▀▀▀     ▀█████████▀  
███   ███   ███    █▄    ███        
███   ███   ███    ███   ███        
 ▀█████▀    ██████████  ▄████▀                             
	`, ansi.Red))
	fmt.Println(ansi.Format("      Yet Another Pokedex\n", ansi.Faint))
}

// REPL logic
func StartRepl() error {

	var config pokeapi.Config

	fmt.Print(ansi.AltScreen)
	fmt.Print(ansi.Clear)

	printBanner()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(ansi.Format("Pokedex > ", ansi.Bold, ansi.Red))
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
