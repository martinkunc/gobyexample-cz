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

// `novaOsoba` vytvoří novou strukturu `osoba` s daným jménem.
func novaOsoba(jmeno string) *osoba {
	// Go je jazyk se správou paměti pomocí garbage
	// collectoru; můžete bezpečně vrátit ukazatel na
	// lokální proměnnou - garbage collector ji uklidí až
	// ve chvíli, kdy na ni neexistují žádné aktivní odkazy.
	o := osoba{jmeno: jmeno}
	o.vek = 42
	return &o
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

	// Je idiomatické zapouzdřit vytváření nových struktur
	// do konstruktorových funkcí.
	fmt.Println(novaOsoba("Jon"))

	// K položkám struktur se přistupuje pomocí tečky.
	s := osoba{jmeno: "Sean", vek: 50}
	fmt.Println(s.jmeno)

	// Struktury je možné měnit.
	s.vek = 51
	fmt.Println(s)

	// Také je možné používat tečku na ukazatel na
	// strukturu - ukazatelé jsou automaticky
	// dereferencovány.
	sp := &s
	sp.vek = 52
	fmt.Println(sp.vek)

	// Pokud typ struktury použijeme jen pro jedinou
	// hodnotu, nemusíme ho pojmenovávat. Hodnota může mít
	// anonymní typ struktury. Tato technika se běžně
	// používá pro [testy řízené tabulkou](testing-and-benchmarking).
	pes := struct {
		jmeno string
		hodny bool
	}{
		"Rex",
		true,
	}
	fmt.Println(pes)
}
