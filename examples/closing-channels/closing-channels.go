// _Zavírání_ kanálu indikuje že do něho nebude poslána
// žádná další zpráva. To se hodí pro oznámení o ukončení
// komunikace příjemcům kanálů.

package main

import "fmt"

// V tomto příkladu použijeme kanál `jobs` pro komunikaci
// práce k provedení z gorutiny `main()` do pracovní
// gorutiny. Když už pro pracovníka nemáme žádnou další
// práci, `zavřeme` kanál `jobs`.
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Tady je pracovní gorutina. Bude opakovaně přijímat
	// z `jobs` pomocí `j, more := <-jobs`. V této
	// speciální 2-hodnotové formě příjmu, hodnota `more`
	// bude `false` když byl `jobs` `zavřen` a všechny
	// hodnoty v kanálu byly už přijaty.
	// Používáme to pro oznámení do `done` když jsme
	// zpracovali všechnu naši práci.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("přijal práci", j)
			} else {
				fmt.Println("přijal všechny práce")
				done <- true
				return
			}
		}
	}()

	// Toto pošle 3 práce pracovníkovi přes kanál `jobs`
	// a poté ho uzavře.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("poslal práci", j)
	}
	close(jobs)
	fmt.Println("poslal všechny práce")

	// Zde čekáme na pracovníka s použitím
	// [synchronizace](channel-synchronization)
	// kterou jsme viděli dříve.
	<-done

	// Čtení ze zavřeného kanálu proběhne okamžitě a vrátí
	// nulovou hodnotu typu, který kanál přenáší.
	// Volitelná druhá návratová hodnota je `true`, pokud
	// přijatá hodnota byla do kanálu doručena úspěšným
	// odesláním, nebo `false`, pokud jde o nulovou hodnotu
	// vrácenou proto, že je kanál zavřený a prázdný.
	_, ok := <-jobs
	fmt.Println("přijal další práce:", ok)
}
