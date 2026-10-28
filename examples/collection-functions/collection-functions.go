// Často v našich programech provádíme operace na kolekcích
// dat, jako je výběr všech prvků splňujících nějaký predikát
// nebo mapování všech prvků do nové kolekce pomocí
// uživatelské kolekce.

// V některých jazycích je idiomatické používat [generické](http://en.wikipedia.org/wiki/Generic_programming)
// datové struktury a algoritmy. Go generika nepodporuje;
// v Go je běžné poskytnout kolekci funkcí jestli a pokud
// jsou specificky potřebné pro váš program a datové typy.

// Zde máme nějaké příklady funkcí pro práci s kolekcemi
// pro výřezy typu `strings`. Můžeš použít tyto příklady
// pro tvorbu svých vlastních funkcí. V některých
// případech může být jasnější prostě vložit kód pro
// manipulaci s kolekcí přímo kam potřebujete,
// místo vytváření a volání pomocné funkce.

package main

import "strings"
import "fmt"

// Index vrací první index cílového řetězce `t`, nebo
// -1 pokud nebyl nalezen žádný záznam.
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include vrací `true` pokud je cílový řetězec t ve
// výřezu.
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any vrací `true` pokud jeden z řetězců ve výřezu
// splňuje predikát `f`.
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All vrací `true` pokud všechny řetězce ve výřezu
// splňují predikát `f`.
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter vrací nový výřez obsahující všechny řetězce
// výřezu které spňují predikát `f`.
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map vrací nový výřez obsahující výsledky po aplikoci
// funkce `f` na každý řetězec v původním výřezu.
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {

	// Tady vyzkoušíme naše různé funkce pro práci
	// s kolekcemi.
	var strs = []string{"broskev", "jablko", "hruška",
		"banán"}

	fmt.Println(Index(strs, "hruška"))

	fmt.Println(Include(strs, "víno"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "b")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "b")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "a")
	}))

	// Všechny příklady výše používaly anonymní funkce,
	// ale také je možné použít pojmenované funkce správného
	// typu.
	fmt.Println(Map(strs, strings.ToUpper))

}
