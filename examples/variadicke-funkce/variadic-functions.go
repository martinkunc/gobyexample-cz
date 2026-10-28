// [_Variadické funkce_](http://en.wikipedia.org/wiki/Variadic_function)
// mohou být volané s libovolným počtem závěrečných
// parametrů
// Například, funkce `fmt.Println` je běžná variadická
// funkce.

package main

import "fmt"

// Zde je funkce která převezme libovolný počet argumentů
// typu `int`.
func sum(cisla ...int) {
	fmt.Print(cisla, " ")
	soucet := 0
	for _, cislo := range cisla {
		soucet += cislo
	}
	fmt.Println(soucet)
}

func main() {

	// Variadické funkce mohou být volané běžným způsobem
	// s jednotlivými argumenty.
	sum(1, 2)
	sum(1, 2, 3)

	// Jestli už máte víc parametrů ve výřezu
	// můžete je aplikovat na variadickou funkci s
	// použitím `func(slice...)`.
	cis := []int{1, 2, 3, 4}
	sum(cis...)
}
