// Příkaz _select_ v Go umožňuje čekat na výsledky
// operací na více kanálech. Kombinování gorutin a kanálů
// a příkazu select je silnou vlastností Go.

package main

import "time"
import "fmt"

func main() {

	// V našem příkladu budeme vybírat pomocí select přes
	// dva kanály.
	c1 := make(chan string)
	c2 := make(chan string)

	// Každý kanál obdrží hodnotu po nějaké době, aby se
	// simulovaly např. blokující RPC operace prováděné v
	// souběžně běžících gorutinách.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "jedna"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "dvě"
	}()

	// Použijeme `select` k čekání na obě z těchto hodnot
	// simultánně, a po přijetí ji vypíšeme.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("přijato", msg1)
		case msg2 := <-c2:
			fmt.Println("přijato", msg2)
		}
	}
}
