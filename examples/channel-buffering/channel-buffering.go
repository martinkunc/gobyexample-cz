// Běžné kanály jsou bez _vyrovnávací paměti_ (bufferu),
// to znamená že budou akceptovat posílání  (`chan <-`)
// pokud existuje korespondující příjemce (`<- chan`) a je
// připravený přijmout poslanou hodnotu.
// _Kanály s vyrovnávací pamětí_ akceptují omezené
// množství hodnot bez korespondujícího příjemce pro tyto
// hodnoty.

package main

import "fmt"

func main() {

	// Zde vytváříme pomocí `make` kanál řetězců s
	// vyrovnávací pamětí pro 2 hodnoty.
	messages := make(chan string, 2)

	// Protože náš kanál má vyrovnávací paměť, můžeme do
	// něj poslat hodnoty bez odpovídajících souběžně
	// probíhajících operací příjmu.
	messages <- "s vyrovnávací pamětí"
	messages <- "kanál"

	// Později můžeme tyto dvě hodnoty přijmout jako
	// obvykle.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
