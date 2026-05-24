package pokeapi

type Config struct {
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
}

type Response struct {
	Config
	Results []Cities `json:"results"`
}

type Cities struct {
	Name string `json:"name"`
}

// Pokemons per area
type PokemonResponse struct {
	Encounters []PokemonEncounters `json:"pokemon_encounters"`
}

type PokemonEncounters struct {
	PokemonRef PokemonRef `json:"pokemon"`
}

type PokemonRef struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Pokemon endpoint
type Pokemon struct {
	Name           string         `json:"name"`
	BaseExperience int            `json:"base_experience"`
	Height         int            `json:"height"`
	Weight         int            `json:"weight"`
	Stats          []PokemonStats `json:"stats"`
	Type           []PokemonType  `json:"types"`
}

type PokemonStats struct {
	StatInfo Stat `json:"stat"`
	BaseStat int  `json:"base_stat"`
}

type Stat struct {
	Name string `json:"name"`
}

type PokemonType struct {
	Type Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
}
