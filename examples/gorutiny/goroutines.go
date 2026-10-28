// _Go rutina_ (goroutine) je lehčí verze běhového vlákna
// nabízeného operačním systémem.

package main

import "fmt"

func f(od string) {
	for i := 0; i < 3; i++ {
		fmt.Println(od, ":", i)
	}
}

func main() {

	// Předpokládejme že máme volání funkce call `f(s)`.
	// Takto bychom ji spustili obvyklým způsobem,
	// kdy poběží synchronně.
	f("přímo")

	// Pro zavolání této funkce v gorutině, použijte
	// `go f(s)`. Tato nová gorutina poběží současně s
	// volající gorutinou (main).
	go f("gorutina")

	// Také je možné nastartovat gorutinu pro anonymní
	// volání funkce.
	go func(msg string) {
		fmt.Println(msg)
	}("jdeme")

	// Naše dvě volání funkcí teď běží asynchronně v
	// oddělených gorutinách, takže provádění propadne až
	// sem. Tato funkce `Scanln` vyžaduje stlačení klávesy
	// než program skončí.
	fmt.Scanln()
	fmt.Println("hotovo")
}
