// Go podporuje
// <a href="http://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>rekurzivní funkce</em></a>.
// Tady je klasický příklad faktoriálu.

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
}
