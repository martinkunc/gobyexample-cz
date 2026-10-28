// Go balíček `sort` implementuje třídění pro zabudované
// a uživatelsky definované typy. Nejdřív se podíváme na
// třídění zabudovaných typů.

package main

import "fmt"
import "sort"

func main() {

	// Mmetody balíčku sort jsou specifické pro zabudované typy;
	// zde je příklad pro řetězce. Všimněte si že třídění je
	// na místě (in-place), takže mění předaný výřez a nevrací
	// nový.
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Retězce:", strs)

	// Příklad třídění `int`ů.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Inty:   ", ints)

	//  Taky můžeme použít `sort` pro kontrolu jestli je výřez
	// už v setříděném pořadí.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Setřízený: ", s)
}
