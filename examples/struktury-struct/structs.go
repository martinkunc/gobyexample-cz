// _Struktury_ v Go jsou kolekcemi prvků s typem.
// Jsou užitečné pro vytváření skupin dat a utváření
// záznamů.

package main

import "fmt"

// Tato struktura `osoba` má položky `jmeno` a `vek`.
type osoba struct {
	jmeno string
	vek   int
}

func main() {

	// Tento zápis vytvoří novou strukturu.
	fmt.Println(osoba{"Bob", 20})

	// Během inicializace struktury je možné pojmenovat
	// položky.
	fmt.Println(osoba{jmeno: "Alice", vek: 30})

	// Vynechané položky budou mít nulovou hodnotu.
	fmt.Println(osoba{jmeno: "Fred"})

	// S prefixem `&` bude vrácený ukazatel na strukturu.
	fmt.Println(&osoba{jmeno: "Ann", vek: 40})

	// K položkám struktur se přistupuje pomocí tečky.
	s := osoba{jmeno: "Sean", vek: 50}
	fmt.Println(s.jmeno)

	// Také je možné používat tečku na ukazatel na
	// strukturu - ukazatelé jsou automaticky
	// dereferencovány.
	sp := &s
	fmt.Println(sp.vek)

	// Struktury je možné měnit.
	sp.vek = 51
	fmt.Println(sp.vek)
}
