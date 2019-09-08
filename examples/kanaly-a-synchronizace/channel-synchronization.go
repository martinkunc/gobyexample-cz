// Kanály můžeme použít pro synchronizaci provádění mezi
// gorutinami. Zde je příklad který s použitím blokujícího
// příjmu čeká až gorutina skončí.
// Až budete čekat na ukončení několika gorutin, možná
// budete preferovat použití [WaitGroup](waitgroup).

package main

import "fmt"
import "time"

// Tohle je funkce která bude čekat na gorutinu. Kanál
// `done` bude použitý k oznámení druhé gorutině že
// práce této funkce je hotova.
func worker(done chan bool) {
	fmt.Print("pracuji...")
	time.Sleep(time.Second)
	fmt.Println("hotovo")

	// Pošle hodnotu abysme oznámili že jsme hotovi.
	done <- true
}

func main() {

	// Nastartuje pracovní gorutinu a předá ji kanál pro
	// oznámení.
	done := make(chan bool, 1)
	go worker(done)

	// Bude blokovat provádění dokud nedostaneme do kanálu
	// oznámení od pracovní rutiny.
	<-done
}
