// V předchozím příkladu jsme viděli jak jednoduše spravovat
// stav counterů pomocí [atomických operací](atomic-counters).
// Pro více komplikované stavy můžeme použít [_mutex_](https://en.wikipedia.org/wiki/Mutual_exclusion)
// pro bezpečný přístup k datům z několika gorutin.

package main

import (
	"fmt"
	"sync"
)

// `Kontejner` obsahuje mapu counterů; protože ho chceme
// měnit souběžně z několika gorutin, přidáme `Mutex`
// pro synchronizaci přístupu.
// Všimněte si, že mutexy se nesmí kopírovat, takže pokud
// tuto strukturu předáváme dál, mělo by to být pomocí
// ukazatele.
type Kontejner struct {
	mu       sync.Mutex
	countery map[string]int
}

func (k *Kontejner) zvys(jmeno string) {
	// Před přístupem ke `countery` zamkneme mutex; na konci
	// funkce ho odemkneme pomocí příkazu [defer](defer).
	k.mu.Lock()
	defer k.mu.Unlock()
	k.countery[jmeno]++
}

func main() {
	k := Kontejner{
		// Všimněte si, že nulová hodnota mutexu je přímo
		// použitelná, takže zde není potřeba žádná
		// inicializace.
		countery: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// Tato funkce v cyklu zvyšuje pojmenovaný counter.
	zvysuj := func(jmeno string, n int) {
		for range n {
			k.zvys(jmeno)
		}
	}

	// Spustíme několik gorutin souběžně; všimněte si,
	// že všechny přistupují ke stejnému `Kontejner`u
	// a dvě z nich ke stejnému counteru.
	wg.Go(func() {
		zvysuj("a", 10000)
	})

	wg.Go(func() {
		zvysuj("a", 10000)
	})

	wg.Go(func() {
		zvysuj("b", 10000)
	})

	// Počkáme, až gorutiny skončí.
	wg.Wait()
	fmt.Println(k.countery)
}
