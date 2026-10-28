// Balíček standardní knihovny `strings` poskytuje mnoho
// užitečných funkcí která pracují s řetězci. Tady jsou některé
// příklady pro získání základního přehledu o balíčku.

package main

import s "strings"
import "fmt"

// Poskytneme `fmt.Println` alias na kratší jméno protože ji budeme
// níže používat často.
var p = fmt.Println

func main() {

	// Zde jsou příklady funkcí dostupných ve
	// `strings`. Protože jsou to funkce v balíčku a ne
	// metody na samotném objektu string,
	// musíme náš řetězec předat funkci jako první parametr.
	// Víc funkcí z balíčku strings můžeš najít
	// v dokumentaci balíčku [`strings`](http://golang.org/pkg/strings/).
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
	p()

	// Není to součástí `strings`, ale stojí za to se zde zmínit o
	// mechanismech jak se získává délka řetězce v bajtech
	// a získání bajtu na daném indexu.
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}

// Povšimněte si, že `len` a indexování výše pracuje na úrovni bajtů.
// Go používá řetězce zakódované v UTF-8, takže často je to užitečné
// jenom takhle. Když budeš pracovat s potencionálně více-bajtovými
// znaky, pravděpodobně budeš chtít použít operace které umí pracovat
// s kódováními. Viz [strings, bytes, runes and characters in Go](https://blog.golang.org/strings)
// pro více informací.
