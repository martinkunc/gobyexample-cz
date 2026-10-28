// Pro čekání na dokončení několika gorutin můžeme použít
// *wait group*.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Toto je funkce, která poběží v každé gorutině.
func worker(id int) {
	fmt.Printf("Worker %d startuje\n", id)

	// Čekáme abychom simulovali dlouhý úkol.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d končí\n", id)
}

func main() {

	// Tato čekací skupina (WaitGroup) je použita pro
	// čekání na ukončení všech gorutin zde
	// nastartovaných. Pozn.: pokud WaitGroup explicitně
	// předáváte do funkcí, mělo by to být *jako ukazatel*.
	var wg sync.WaitGroup

	// Nastartujeme několik gorutin pomocí `WaitGroup.Go`.
	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			worker(i)
		})
	}

	// Blokujeme, dokud nejsou hotové všechny gorutiny
	// nastartované pomocí `wg`. Gorutina je hotová, když
	// se funkce, kterou volá, vrátí.
	wg.Wait()

	// Všimněte si, že tento přístup nenabízí přímočarý
	// způsob, jak předat chyby z pracovních rutin. Pro
	// pokročilejší případy použití zvažte
	// [balíček errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup).
}
