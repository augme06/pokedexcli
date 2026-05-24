package repl

import (
	"testing"

	"github.com/augme06/pokedexcli/internal/pokeapi"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "multiple spaces",
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "uppercase with accents",
			input:    "Olá amigos",
			expected: []string{"olá", "amigos"},
		},
		{
			name:     "mixed case with trailing spaces",
			input:    "hey Hey  ",
			expected: []string{"hey", "hey"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "only whitespace",
			input:    "   \t  \n  ",
			expected: []string{},
		},
		{
			name:     "single word",
			input:    "PIKACHU",
			expected: []string{"pikachu"},
		},
		{
			name:     "numbers and special chars",
			input:    "test123 PokéAPI!",
			expected: []string{"test123", "pokéapi!"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := pokeapi.CleanInput(tc.input)

			if len(actual) != len(tc.expected) {
				t.Fatalf("input %q: got %d words, expected %d; result: %v", tc.input, len(actual), len(tc.expected), actual)
			}

			for i, word := range actual {
				if word != tc.expected[i] {
					t.Errorf("word %d: got %q, expected %q", i, word, tc.expected[i])
				}
			}
		})
	}
}
