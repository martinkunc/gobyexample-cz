// _Go rutina_ (goroutine) je lehčí verze běhového vlákna
// nabízeného operačním systémem.

package main

import (
	"fmt"
	"time"
)

func f(od string) {
	for i := range 3 {
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
	// oddělených gorutinách. Počkáme, až skončí (pro
	// spolehlivější přístup použijte [WaitGroup](waitgroups)).
	time.Sleep(time.Second)
	fmt.Println("hotovo")
}
