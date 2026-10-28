// Někdy budeme chtít třídit kolekci podle něčeho jiného
// než je její přirozené řazení. Například, předpokládejme
// že jsme chtěli setřídit řetězce podle jejich délky
// a ne abecedně. Zde je příklad volitelných třídění
// v Go.

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// Naimplementujeme porovnávací funkci pro délky
	// řetězců. Hodí se nám k tomu `cmp.Compare`.
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	// Teď můžeme zavolat `slices.SortFunc` s touto vlastní
	// porovnávací funkcí a setřídit `fruits` podle délky
	// názvu.
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	// Stejnou techniku můžeme použít i pro třídění výřezu
	// hodnot, které nejsou zabudovanými typy.
	type Osoba struct {
		jmeno string
		vek   int
	}

	lide := []Osoba{
		Osoba{jmeno: "Jax", vek: 37},
		Osoba{jmeno: "TJ", vek: 25},
		Osoba{jmeno: "Alex", vek: 72},
	}

	// Setřídíme `lide` podle věku pomocí `slices.SortFunc`.
	//
	// Pozn.: pokud je struktura `Osoba` velká, možná
	// budete chtít, aby výřez obsahoval `*Osoba`,
	// a třídicí funkci tomu přizpůsobíte. Pokud si
	// nejste jistí, [změřte to](testing-and-benchmarking)!
	slices.SortFunc(lide,
		func(a, b Osoba) int {
			return cmp.Compare(a.vek, b.vek)
		})
	fmt.Println(lide)
}
