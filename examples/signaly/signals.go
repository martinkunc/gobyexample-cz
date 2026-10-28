// Někdy bychom chtěli aby naše Go programy inteligentně
// ošetřovali [Unixové signály](http://en.wikipedia.org/wiki/Unix_signal).
// Například, můžeme chtít aby server we might want a server to elegantně
// skončil když dostaneme signál `SIGTERM`, nebo aby nástroj
// příkazové řádky přestal zpracovávat vstup pokud obdrží `SIGINT`.
// Zde je jak ošetřit signály v Go pomocí kanálů.

package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {

	// Go notifikace o signálech pracuje tak, že posílá
	// hodnoty `os.Signal` do kanálu. Vytvoříme kanál pro
	// příjem takových notifikací (a také jeden kde se sami
	// notifikujeme že program může skončit).
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// `signal.Notify` registruje daný kanál pro příjem notifikací
	// o specifikovaných signálech.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Tato goroutina provádí blokující příjem signálů.
	// Když nějaký dostané tak ho vypíše a oznámí programu
	// že může skončit.
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// Program zde bude čekat dokud nedostane
	// očekávaný signál (jak je naznačeno v gorutině výše
	// posílající hodnotu když je `done`) a pak skončí.
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
