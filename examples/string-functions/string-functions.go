// Balíček standardní knihovny `strings` poskytuje mnoho
// užitečných funkcí která pracují s řetězci. Tady jsou některé
// příklady pro získání základního přehledu o balíčku.

package main

import (
	"fmt"
	s "strings"
)

// Poskytneme `fmt.Println` alias na kratší jméno protože ji budeme
// níže používat často.
var p = fmt.Println

func main() {

	// Zde jsou příklady funkcí dostupných ve
	// `strings`. Protože jsou to funkce v balíčku a ne
	// metody na samotném objektu string,
	// musíme náš řetězec předat funkci jako první parametr.
	// Víc funkcí najdete v dokumentaci balíčku
	// [`strings`](https://pkg.go.dev/strings).
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}
