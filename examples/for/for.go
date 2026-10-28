// `for` je v Go jediná konstrukce pro cykly. Zde jsou
// tři základní typy `for` cyklů.

package main

import "fmt"

func main() {

	// Nejzákladnější cyklus s jednoduchou podmínkou
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Klasický cyklus ve tvaru
	// inicializace/podmínka/popodmínkový příkaz
	// `for` cyklu.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// `for` bez podmínky bude cyklit opakovaně
	// dokud ho nepřerušíte voláním `break` a nevyskočíte
	// z cyklu nebo voláním `return` nevyskočíte z funkce
	// kde se příkaz nachází
	for {
		fmt.Println("smyčka")
		break
	}

	// Také můžete přeskočit na provádění další iterace
	// cyklu použitím příkazu `continue`
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
