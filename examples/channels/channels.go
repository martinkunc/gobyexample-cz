// _Kanály_ jsou potrubí které propojuje souběžné
// gorutiny. Do kanálů můžete posílat hodnoty z jedné
// gorutiny a přijímat tyto hodnoty v jiné gorutině.

package main

import "fmt"

func main() {

	// Nový kanál vytvoříte s pomocí
	// `make(chan typ-hodnot)`. Kanály mají typy podle
	// hodnot, jaké doručují.
	messages := make(chan string)

	// Hodnotu do kanálu _pošlete_ pomocí zápisu
	// `kanál <-`. Zde posíláme z nové gorutiny hodnotu
	// `"ping"`do kanálu `messages` který jsme vytvořili
	// výše.
	go func() { messages <- "ping" }()

	// Zápis `<-kanál` _přijímá_ hodnotu z kanálu.
	// Zde dostaneme zprávu `"ping"` kterou jsme poslali
	// výše a vypíšeme ji.
	msg := <-messages
	fmt.Println(msg)
}
