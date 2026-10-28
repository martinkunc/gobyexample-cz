// [Časovače](casovace=timers) jsou pro situace kdy chcete
// něco udělat jednou v budoucnu - opakující časovače
// _tickers_ jsou pro situace kdy chcete něco provádět
// opakovaně v pravidelném interalu.
// Tady je příkad opakujícího se časovače který tiká
// periodicky dokud ho nezastavíme.

package main

import "time"
import "fmt"

func main() {

	// Opakující časovače používají podobný mechanismus
	// jako běžné časovače: kanál kterému je posílaná
	// hodnota. Zde použijeme zabudovaný příkaz `range`
	// pro iterování přes hodnoty které do kanálu
	// přijdou každých 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tik v", t)
		}
	}()

	// Opakující časovače je možné stejně jako běžné
	// časovače zastavit. Jakmile je časovač zastavený
	// už nedostane do kanálu žádnou další hodnotu.
	// Náš zastavíme po 1600ms.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker zastaven")
}
