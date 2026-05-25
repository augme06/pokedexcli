package ansi

import "strings"

const (
	Esc = "\x1b[" // Escape
)

const (
	AltScreen     = Esc + "?1049h"
	ExitAltScreen = Esc + "?1049l"
	Clear         = Esc + "H" + Esc + "2J"
)

const (
	Reset  = Esc + "0m" // Prevents all the text from being colored
	Bold   = Esc + "1m"
	Faint  = Esc + "2m"
	Red    = Esc + "31m"
	Green  = Esc + "32m"
	Yellow = Esc + "33m"
	Blue   = Esc + "34m"
)

func Format(text string, codes ...string) string {
	var prefix strings.Builder
	for _, code := range codes {
		prefix.WriteString(code)
	}
	return prefix.String() + text + Reset
}
