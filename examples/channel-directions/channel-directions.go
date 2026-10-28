// Když používáte kanály jako parametry funkcí, můžete
// specifikovat jestli je kanál zamýšlený jen pro posílání
// nebo příjem hodnot. Tohle konkrétně zvyšuje typovou
// bezpečnost programu.

package main

import "fmt"

// Tato funkce `ping` akceptuje jen kanál pro posílání
// hodnot. Pokud byste zkusili přijímat na tomto kanálu,
// způsobilo by to chybu při kompilaci.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Funkce `pong` přijímá jeden kanál pro příjem
// (`pings`) a druhý pro posílání (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "posílaná zpráva")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
