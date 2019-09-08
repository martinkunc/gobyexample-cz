// V tomto příkladu se podíváme jak naimplementovat
// _zásobníky pracovních rutin_ (worker pools) s pomocí
// gorutin a kanálů.

package main

import "fmt"
import "time"

// Tady je pracovní rutina které spustíme několik instancí
// soubežně. Tyto rutiny přijmou práce na kanálu `jobs` a
// pošlou odpovídající výsledky do `results`.
// V každém jobu budeme čekat vteřinu, abychom simulovali
// náročný úkol.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "nastartoval job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "ukončil job", j)
		results <- j * 2
	}
}

func main() {

	// Pro náš zásobník rutin jim budeme chtít poslat
	// práci a pak přijmout a poté posbírat jejich
	// výsledky. K tomu si vytvoříme 2 kanály.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Tohle nastartuje 3 pracovní rutiny, nejdřív
	// zablokované, protože ještě nemají žádnou práci
	// v kanálu jobs.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Zde pošleme 5 `job`ů a potom kanál zavřeme s
	// `close` pro oznámení že to byla všechna práce
	// kterou máme.
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Na závěr posbíráme všechny výsledky práce.
	// To také zajistí že pracovní gorutiny skončily.
	// Alternativně můžeme také počkat na víc gorutin s
	// použitím [Čekání na skupiny s WaitGroup](cekani-na-skupiny-s-waitgroup).
	for a := 1; a <= 5; a++ {
		<-results
	}
}
