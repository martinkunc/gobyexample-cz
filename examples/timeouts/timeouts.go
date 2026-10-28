// _Timeouty_ (časová omezení) jsou důležitá pro programy
// připojující se k externím zdrojům a nebo ty co jakkoliv
// potřebují omezovat čas provádění. Implementace timeoutů
// je v Go díky kanálům a `select`ům elegantní.

package main

import "time"
import "fmt"

func main() {

	// V našem příkladě předpokládejme že provádímě
	// externí volání které vrací svůj výsledek na kanálu
	// `c1` po 2s.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "výsledek 1"
	}()

	// Zde je `select` implementující timeout.
	// `res := <-c1` čeká na výsledek a `<-Time.After`
	// očekává obdrženou hodnotu po timeoutu 1s.
	// A protože `select` pokračuje s první přijatou
	// hodnotou která je připravena, vybere si větev
	// timeoutu pokud bude operace trvat více než
	// povolenou 1s.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// Pokud dovoleníme delší timeout 3s, potom uspěje
	// příjem z `c2` a my vypíšeme výsledek.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "výsledek 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
