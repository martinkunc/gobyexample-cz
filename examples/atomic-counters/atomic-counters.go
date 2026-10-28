// Primární mechanismus pro správu stavů je v Go
// komunikace pomocí kanálů. To jsme viděli v příkladu se
// [zásobníky pracovních rutin](worker-pools). Ale existují ještě
// další možnosti pro správu stavu. Tady se podíváme
// na použití _atomických counterů_ z balíčku `sync/atomic`
// které jsou používané zároveň několika gorutinami.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// Použijeme atomický celočíselný typ pro reprezentaci
	// našeho (vždy kladného) counteru.
	var ops atomic.Uint64

	// WaitGroup nám pomůže počkat, až všechny gorutiny
	// dokončí svou práci.
	var wg sync.WaitGroup

	// Nastartujeme 50 gorutin, z nichž každá zvýší
	// counter přesně 1000 krát.
	for range 50 {
		wg.Go(func() {
			for range 1000 {
				// Pro atomické zvýšení counteru použijeme `Add`.
				ops.Add(1)
			}
		})
	}

	// Počkáme, až budou všechny gorutiny hotové.
	wg.Wait()

	// Tady už do `ops` žádná gorutina nezapisuje, ale
	// pomocí `Load` je bezpečné atomicky číst hodnotu
	// i ve chvíli, kdy ji jiné gorutiny (atomicky) mění.
	fmt.Println("ops:", ops.Load())
}
