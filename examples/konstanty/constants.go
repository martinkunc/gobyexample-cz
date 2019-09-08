// Go podporuje _konstanty_ znakové, řetězcové,
// booleanovské, a číselné hodnoty.

package main

import "fmt"
import "math"

// `const` deklaruje konstantní hodnotu.
const s string = "konstanta"

func main() {
	fmt.Println(s)

	// Příkaz `const` jde použít kdekoliv kde je možné
	// použít příkaz `var`
	const n = 500000000

	// Konstantní výrazy mohou provádět aritmetiku s
	// libovolnou přeností.
	const d = 3e20 / n
	fmt.Println(d)

	// Číselná konstanta nemá žádný typ, dokud ji nějaký
	// není daný, jako například explicitní konverzí.
	fmt.Println(int64(d))

	// Číselná konstanta získá svůj typ až ve chvíli kdy
	// se použije na místě, které typový parametr
	// vyžaduje. Takové místo je třeba přiřazení a nebo
	// zavolání funkce. Například funkce
	// `math.Sin` očekává hodnotu typu `float64`.
	fmt.Println(math.Sin(n))
}
