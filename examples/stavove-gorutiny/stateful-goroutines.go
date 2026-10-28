// V předchozím příkladu jsme použili explicitní zamykání s
// [mutexy](mutexy) pro synchronizaci přístupu ke sdílenému stavu
// přes několik gorutin. Jinou možností k dosažení toho stejného
// výsledku je použít zabudované vlastnosti pro synchronizaci
// goroutin a kanálů. TEnto přístup založený na kanálech
// odpovídá myšlenkám Go se sdílením paměti pomocí
// komunikace, kdy každý kousek dat je vlastněný
// přesně 1 goroutinou.

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// V tomto příkladu bude náš stav vlastněný jednou
// gorutinou. To zajistí že data nebudou nikdy během souběžného
// přístupu poškozena. K přečtení nebo zápisu tohoto stavu
// ostatní gorutiny pošlou zprávu vlastnící gorutině
// a obdrží odpovídající odpovědi.
// Tyto struktury `readOp` a `writeOp` zabalují tento požadavek
// a způsob jakým vlastnící gorutina odpoví.
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// Stejně jako předtím budeme počítat kolik operací jsme udělali.
	var readOps uint64
	var writeOps uint64

	// Kanály `reads` a `writes` budou ostatními gorutinami
	// využity pro zavolání požadavku na zápis nebo čtení.
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Zde je gorutina která vlastní `stav` který je stejně
	// jako v předchozím příkladu mapa, ale nyní je privátní
	// jen pro stavovou gorutinu. Tato gorutina opakovaně
	// vybírá za kanálů `reads` aa `writes` a odpovídá na
	// požadavky když přijdou. Jako odpověď se nejprve provede
	// požadovaná operace a poté se pošle hodnota na
	// odpovídající kanál `resp` pro oznámení úspěchu
	// (a v případě `reads` také požadovaná hodnota).
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Toto nastaruje 100 gorutin pro vyvolání čtení od
	// stav vlastnící gorutiny skrze kanál `reads`.
	// Každé čtení vyžaduje vytvoření `readOp` a jejího
	// odeslání skrz kanál `reads` a poté přijetí odpovědi
	// v poskytnutém kanálu `resp`.
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// A také nastartujeme 10 zápisů s použitím stejného
	// přístupu.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Nechá gorutiny vteřinu pracovat.
	time.Sleep(time.Second)

	// Na závěr načti a vypiš počty operací.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
