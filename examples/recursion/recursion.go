// Go podporuje
// <a href="https://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>rekurzivní funkce</em></a>.
// Tady je klasický příklad.

package main

import "fmt"

// Tato funkce `fact` volá samu sebe dokud nedosáhne
// základního stavu `fact(0)`.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// Rekurzivní mohou být i anonymní funkce, ale to
	// vyžaduje explicitně deklarovat proměnnou pomocí `var`,
	// do které funkci uložíme, ještě před její definicí.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// Protože jsme `fib` předtím deklarovali v `main`,
		// Go ví, kterou funkci má zde pomocí `fib` zavolat.
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
