// Pro čekání na dokončení několika gorutin můžeme použít
// *wait group*.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Toto je funkce která poběží v každé gorutině
// Všimněte si že WaitGroup musí být předána funkcím jako
// ukazatel.
func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d startuje\n", id)

	// Čekáme abychom simulovali dlouhý úkol.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d končí\n", id)

	// Oznamujeme do WaitGroup že tato rutina je hotova.
	wg.Done()
}

func main() {

	// Tato čekací skupina (WaitGroup) je použita pro
	// čekání na ukončení všech gorutin zde
	// nastartovaných
	var wg sync.WaitGroup

	// Nastartujeme několik gorutin a pro každou zvýšíme
	// počítadlo ve WaitGroup.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Blokujeme dokud počítadlo ve WaitGroup nespadne
	// zpět na 0; Tj že všechny rutiny oznámily že jsou
	// hotové.
	wg.Wait()
}
