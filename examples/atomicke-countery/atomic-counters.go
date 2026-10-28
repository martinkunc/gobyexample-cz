// Primární mechanismus pro správu stavů je v Go
// komunikace pomocí kanálů. To jsme viděli v příkladu se
// [zásobníky pracovních rutin](zasobniky-pracovnich-rutin-worker-pools). Ale existují ještě
// další možnosti pro správu stavu. Tady se podíváme
// na použití _atomických counterů_ z balíčku `sync/atomic`
// které jsou používané zároveň několika go rutinami.

package main

import "fmt"
import "time"
import "sync/atomic"

func main() {

	// Použijeme bezznaménkový integer pro reprezentaci našeho vždy kladného
	// counteru.
	var ops uint64

	// Pro simulaci souběžně probíhajících změn nastartujeme 50
	// goroutin. Každá z nich zvýší counter jednou za milisekundu.
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// Pro atomické zvýšení counteru použijeme
				// `AddUint64` a předáme ji adresu v paměti
				// našeho counteru `ops` s použitím  `&` syntaxe.
				atomic.AddUint64(&ops, 1)

				// Trochu počkej mezi zvýšeními.
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Počkej vteřinu aby se nahromadily nějaké operace.
	time.Sleep(time.Second)

	// Abychom bezpečne získali hodnotu counteru zatímco
	// je stále měněn ostatními gorutinami, extrahujeme
	// kopii jeho aktuální hodnoty do `opsFinal` s použitím
	// `LoadUint64`. Stejně jako výše potřebujeme dát funkci
	// adresu v paměti `&ops` ze které by se měla hodnota
	// načíst
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
