// Go má zabudovanou podporu pro
// _více návratových hodnot_.
// Tato vlastnost je často je v idiomatickém Go používána,
// například pro vrácení obou hodnot, výsledku i chyby
// z funkce.

package main

import "fmt"

// Výraz `(int, int)` v této signatuře funkce ukazuje že
// funkce vrátí 2 `int`egery.
func hodnoty() (int, int) {
	return 3, 7
}

func main() {

	// Zde používáme dvě různé návratové hodnoty z volání
	// pomocí _vícenásobného přiřazení_.
	a, b := hodnoty()
	fmt.Println(a)
	fmt.Println(b)

	// Pokud potřebujete jenom podmnožinu z vrácených
	// hodnot, použijte zahazovací identifikátor `_`.
	_, c := hodnoty()
	fmt.Println(c)
}
