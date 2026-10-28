// Go balíček `slices` implementuje třídění pro zabudované
// a uživatelsky definované typy. Nejdřív se podíváme na
// třídění zabudovaných typů.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// Třídicí funkce jsou generické a fungují pro jakýkoli
	// _uspořádaný_ zabudovaný typ. Seznam uspořádaných
	// typů najdete v [cmp.Ordered](https://pkg.go.dev/cmp#Ordered).
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Řetězce:", strs)

	// Příklad třídění `int`ů.
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Inty:   ", ints)

	// Balíček `slices` můžeme také použít pro kontrolu,
	// jestli je výřez už v setříděném pořadí.
	s := slices.IsSorted(ints)
	fmt.Println("Setříděný: ", s)
}
