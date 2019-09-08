// V předchozím příkladu jsme viděli jak jednoduše spravovat
// stav counterů pomocí [atomických counterů](atomicke-countery).
// Pro více komplikované stavy můžeme použít <em>[mutexy](http://en.wikipedia.org/wiki/Mutual_exclusion)</em>
// pro bezpečný přístup k datům z několika goroutin.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	// Pro náš příklad bude `stavem` mapa.
	var state = make(map[int]int)

	// Tento `mutex` bude ke `stavu` synchronizovat přístup.
	var mutex = &sync.Mutex{}

	// Budeme sledovat kolik čtecích a zapisovacích operací
	// uděláme.
	var readOps uint64
	var writeOps uint64

	// Zde nastartujeme 100 gorutin pro provádění opakovaných
	// čtení ze stavu, jednou za milisekundu v každé gorutině.
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {

				// Pro každé čtení vybereme klíč pro přístup,
				// pomocí `Lock()` zamkneme `mutex` k zajištění
				// exkluzivního přístupu ke `stavu`, načteme
				// hodnotu se zvoleným klíčem,
				// pomocí `Unlock()` odemkneme mutex, a zvýšíme
				// počet v `readOps`.
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				// Trochu počkejme mezi čteními.
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Také nastartujeme 10 gorutin pro simulaci zápisů
	// s použitím stejného přístupu jako jsme použili u čtení.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Nechejme 10 gorutin pracovat na `stavu` a
	// `mutexu` po jednu vteřinu.
	time.Sleep(time.Second)

	// Načti a vypiš výsledné počty operací.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// Se závěrečným zamknutím `stavu`, ukaž jak jsme skončili.
	mutex.Lock()
	fmt.Println("stav:", state)
	mutex.Unlock()
}
