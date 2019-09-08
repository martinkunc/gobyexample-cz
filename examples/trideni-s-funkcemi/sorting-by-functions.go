// Někdz budeme chtít třídit kolekci podle něčeho jiného
//  než je její přirozené řazení. Například, předpokládejme
// že jsme chtěli setřídit řetězce podle jejich délky
// a ne abecedně. Zde je příklad volitelných třídění
// v Go.

package main

import "sort"
import "fmt"

// Abychom v Go setřídili podle zvolené funkce, budeme
// potřebovat odpovídající typ. Vytvořili jsme tady typ
// `byLength` který je jenom alias na zabudovaný typ
// `[]string`.
type byLength []string

// Implementujeme interface `sort.Interface` - `Len`, `Less`, a
// `Swap` - na našem typu, takže můžeme použít generickou funkci
// `Sort` z balíčku `sort. `Len` a `Swap` budou
// obvykle podobné napříč typy a `Less` bude obsahovat
// vlastní třídící logiku. V našem případě chceme třídit podle
//  délky řetězce vzestupně, takže zde použijeme
// `len(s[i])` a `len(s[j])`.
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// Když máme všechno připravené, nyní můžeme naimplementovat
// naše volitelné třídění překonvertováním původního výřezu `fruits`
// na `byLength` a poté na výřez s tímto typem použít  `sort.Sort`.
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
